# gRPC examples

## Build Container

```
make image
```

## Run Container: This would need to be run in Google Clour Run to make use of the Firestore

Push the container image up to Google Container Registry

Then you can visit `https://<CONTAINER_URL>/v1/status` to see it running.

Alternatively, use curl:

```
curl -X POST -H 'Content-type: application/json' --data '{"name": "Electric Boogaloo", "height": 15, "width": 30, "depth": 20}' https://<CONTAINER_URL>/v1/make-box
curl -X GET "https://<CONTAINER_URL>/v1/boxes" -H "accept: application/json"
curl -X GET "https://<CONTAINER_URL>/v1/status" -H "accept: application/json"
```