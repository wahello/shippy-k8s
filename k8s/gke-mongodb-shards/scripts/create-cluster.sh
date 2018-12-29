project="cgault-sandbox"
zone="us-central1-a"
cluster_name="shippy"
cluster_version="1.11.5-gke.5"
machine_type="n1-standard-1"
num_nodes="3"

gcloud beta container --project ${project} clusters create ${cluster_name} --zone ${zone} --cluster-version ${cluster_version} --machine-type ${machine_type} --num-nodes ${num_nodes} --enable-cloud-logging --enable-cloud-monitoring --enable-autoupgrade --enable-autorepair

# gcloud beta container --project ${project} clusters delete ${cluster_name} --zone ${zone}