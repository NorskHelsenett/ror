package auditservice

import (
	"context"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/cmd/audit/helsegitlabclient"
	aclrepository "github.com/NorskHelsenett/ror/internal/acl/repositories"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func init() {
	rlog.Debugc(context.Background(), "Init audit service")
}

func CreateAndCommitAclList(ctx context.Context, event messagebuscontracts.AclUpdateEvent) {
	acls, err := aclrepository.GetAllACL2(ctx)
	if err != nil {
		rlog.Fatalc(ctx, "could not get acl items ...", nil)
	}

	md, err := createMarkdown(ctx, acls)
	if err != nil {
		rlog.Fatalc(ctx, "could not create markdown of acl list ...", nil)
	}

	err = helsegitlabclient.PushAclToRepo(md)
	if err != nil {
		rlog.Fatalc(ctx, "could not push markdown to repo ...", err)
	}

	rlog.Debugc(ctx, "acl updated")
}

func createMarkdown(ctx context.Context, acls []aclmodels.AclV2ListItem) (string, error) {
	var sb strings.Builder
	// not indenting because of result in file
	sb.WriteString(`# Rolle og rettigheter

## Liste

| Group | Scope | Subject | Read | Create | Update | Delete | Owner | Kubernetes.Logon | Issued by | Created |
|---|---|---|---|---|---|---|---|---|---|---|
`)

	for _, acl := range acls {
		aclItem := fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			acl.Group, string(acl.Scope), acl.Subject,
			getEmojiByBool(acl.Access.Read), getEmojiByBool(acl.Access.Create), getEmojiByBool(acl.Access.Update), getEmojiByBool(acl.Access.Delete), getEmojiByBool(acl.Access.Owner),
			getEmojiByBool(acl.Kubernetes.Logon), acl.IssuedBy, acl.Created.String())
		sb.WriteString(aclItem)
	}

	md := sb.String()
	return md, nil
}

func getEmojiByBool(boolean bool) string {
	if boolean {
		return ":white_check_mark:"
	} else {
		return ":x:"
	}
}
