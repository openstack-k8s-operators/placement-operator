cat > ci/olm.yaml <<EOF_CAT
---
apiVersion: v1
kind: Namespace
metadata:
    name: openstack-operators
    labels:
      pod-security.kubernetes.io/enforce: privileged
      security.openshift.io/scc.podSecurityLabelSync: "false"
---
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: placement-operator-index
  namespace: openstack-operators
spec:
  image: ${CATALOG_IMG}
  sourceType: grpc
---
apiVersion: operators.coreos.com/v1
kind: OperatorGroup
metadata:
  name: openstack
  namespace: openstack-operators
---
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: placement-operator
  namespace: openstack-operators
spec:
  name: placement-operator
  channel: alpha
  source: placement-operator-index
  sourceNamespace: openstack-operators
EOF_CAT
