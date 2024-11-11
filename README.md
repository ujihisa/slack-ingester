# Slack Ingester

```
Slack -> Google Cloud Function -> Google Cloud Pub/Sub -> Google Cloud Run -> Slack
         ^^^^^^^^^^^^^^^^^^^^^^^^^
         This is Slack Ingester

```

## Upgrade dependency libraries

```bash
go get -u
go mod tidy
```

## Deploy

manual

```
git push
/opt/google-cloud-cli/bin/gcloud functions deploy slack-ingester --source https://source.developers.google.com/projects/devs-sandbox/repos/github_ujihisa_slack-ingester/moveable-aliases/master/paths/ --runtime go122
```

## LICENCE

GPL version 3 or any later versions
