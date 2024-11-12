// Read this first:
// https://www.mongodb.com/docs/mongodb-shell/write-scripts/#execute-a-script-from-within-mongosh

print('Starting migration')

use('nhn-ror')

db.clusters.find().forEach(function (cluster) {
  print('id: ' + cluster.clusterid)

  // update clusters collection
  db.clusters.updateMany(
    { clusterid: cluster.clusterid },
    {
      $set: {
        clusteridold: cluster.clusterid,
        clusterid: cluster.identifier,
      },
    }
  )

  // update acl collection
  db.acl.updateMany(
    {
      subject: cluster.clusterid,
      scope: 'cluster',
    },
    {
      $set: {
        subject: cluster.identifier,
      },
    }
  )

  // update apikeys collection
  db.apikeys.updateMany(
    {
      identifier: cluster.clusterid,
      type: 'Cluster',
    },
    {
      $set: {
        identifier: cluster.identifier,
        displayname: cluster.identifier,
      },
    }
  )

  // update metrics collection
  db.metrics.updateMany(
    { 'metadata.clusterId': cluster.clusterid },
    { $set: { 'metadata.clusterId': cluster.identifier } }
  )

  // TODO: audit logs?!
  // TODO: switchboard?!
})

db.resources.deleteMany({})

print('Migration done')
