# Patching resources

A `Resource` carries a whole resource. Replacing it is fine when the sender
owns everything it sends, which is how agents report. It does not work when a
caller wants to change one field and leave the rest alone, because the sender
would have to send back everything it did not want to change, and the two
senders would overwrite each other.

`ResourcePatch` is the partial update counterpart. It carries only the fields
the caller named.

## The three states

A partial update needs to tell three things apart:

| State | Meaning | JSON |
| --- | --- | --- |
| unset | leave the stored field as it is | field absent |
| null | reset the stored field to its zero value | `"field": null` |
| value | write the value to the stored field | `"field": "value"` |

A plain Go value only carries two of them: after decoding, a field that was
absent and a field that was sent as its zero value look identical. Every field
of a patch is therefore an `optional.Optional[T]`
(`pkg/helpers/optional`), which carries all three.

Every optional field is tagged `json:",omitzero"`. Without it an unset field
would be encoded as `null`, and a round trip through JSON would silently turn
"untouched" into "cleared".

## What a patch looks like

```go
patch := rorresources.NewResourcePatchKubernetesCluster(uid).
    SetLabel("team", "platform").
    RemoveTag("deprecated").
    SetKubernetesCluster(rortypes.ResourcePatchKubernetesCluster{
        Spec: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterSpec{
            ClusterMetadata: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
                Description: optional.NewOptional("a new description"),
                Contacts:    optional.Null[[]rortypes.ResourceContact](),
            }),
        }),
    })

patched, err := patch.Apply(current)
```

`Apply` returns a copy; the resource passed in is never modified. The copy is
rebuilt with `NewResourceFromStruct`, so `ApplyInputFilter` and `GenRorHash`
work on the result like on any other resource. The expected server flow is
therefore read, apply, write.

## The two parts of a patch

A patch carries two independent parts.

The **common part** is embedded in every patch and works on every kind. It
covers labels, annotations and the ror metadata including tags. A patch that
only touches the common part is valid for any kind, which is what makes
tagging a resource possible without a mirror for its kind.

The **sub resource part** covers the resource definition itself. It only
exists for the kinds marked `Patchable: true` in `rordefs`, because those are
the kinds with a generated patch mirror.

Not everything in the common part is patchable, by design:

- Of `metav1.ObjectMeta` only labels and annotations are. The rest is identity
  and bookkeeping that a partial update must not rewrite.
- `RorMeta.Hash` is not, because `GenRorHash` derives it from the resource.
  Patching it would only let a caller store a hash that does not match.
- `TypeMeta` is not. The kind and apiVersion of a patch identify its target,
  they are not a payload.

## How collections are patched

Maps are patched **per key**, through `optional.Map`:

```json
{"labels": {"team": "platform", "stale": null}}
```

sets `team` and removes `stale`, leaving every other label alone. Setting the
field itself to `null` clears the whole map.

Tags are stored as a slice but behave like a map, so they are patched by tag
key through `ResourcePatchTags`. An updated tag keeps its position, a new tag
is appended.

Slices are replaced **as a whole**, because a slice has no stable key to
address an element by. The same goes for types from other packages and types
with custom JSON marshalling: they are not ours to take apart.

## Making a kind patchable

1. Set `Patchable: true` on the kind in `pkg/rorresources/rordefs/defs.go`.
2. Run the generator.

That is the only place to change. The generator reads the resource definition
off the `<Kind>Resource` field that `rorresources.Resource` already carries for
the kind, which is generated from the same `rordefs` entry and is what the
patch is applied to, so there is nothing to keep in step by hand.

That produces the mirror types in
`pkg/rorresources/rortypes/resourcepatchdef_<definition>.go` and the field,
constructor, setter and apply case in `pkg/rorresources/resourcepatch.go`.
Neither file is edited by hand.

A mirror is written to the file matching the one its source type is declared
in, so `resourcedef_kubernetescluster.go` gets a
`resourcepatchdef_kubernetescluster.go` beside it. Following the declaration
rather than the kind matters because resource definitions share types: the
kubernetescluster definition is nested in both workspace and datacenter. A
shared type is declared once, in the file it belongs to, so making another
kind patchable adds files instead of reshuffling the ones that exist.

Mirrors left over from an earlier run are removed, so a kind that stops being
patchable does not leave a stale file behind. A `resourcepatchdef_` file
without the generated header is never touched, which is what keeps the hand
written `resourcepatchdef_common.go` safe.

Generating the mirrors rather than writing them is what keeps them honest: a
field added to a resource definition becomes patchable in the same commit,
instead of drifting until someone notices.

If a struct in a definition should be replaced as a whole rather than taken
apart, add its name to `Opaque` in the generator config. If a field should not
be patchable at all, add `"TypeName.FieldName"` to `Skip`.
