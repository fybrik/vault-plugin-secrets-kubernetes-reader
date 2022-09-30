kubectl create secret generic dummy-secret -n default \
  --from-literal=username=devuser \
  --from-literal=password='S!B\*d$zDsb='
