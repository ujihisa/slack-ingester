# Slack Ingester

```
Slack -> Google Cloud Function -> Google Cloud Pub/Sub -> Google Cloud Run -> Slack
         ^^^^^^^^^^^^^^^^^^^^^^^^^
         This is Slack Ingester

```

## Deploy

manual

```
git push
gcloud functions deploy slack-ingester --source https://source.developers.google.com/projects/devs-sandbox/repos/github_ujihisa_slack-ingester/moveable-aliases/master/paths/
```

## LICENCE

GPL version 3 or any later versions
