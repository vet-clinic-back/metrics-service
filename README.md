# metrics-service

**ENV** vars you can see [here](.env.example). And check it [here](https://github.com/vet-clinic-back/project-setup-props/blob/main/.env).

## INPUT TCP DTO

```json lines
{
  "id": 1,
  "device_id": 12345,
  "pulse": 72.5,
  "temperature": 36.6,
  "LoadCell": {
    "output1": 10.5,
    "output2": 12.3
  },
  "MuscleActivity": {
    "output1": 5.8,
    "output2": 6.1
  },
  "timestamp": "2025-02-24T12:34:56Z" // i will
}

```

## TODO
- [ ] Change metrics time.
- [X] Return metrics on get request
- [X] Receive metrics
- [X] Put metrics to database
- [X] Swagger

## Error resolve
- https://stackoverflow.com/questions/33893150/dial-tcp-lookup-xxx-xxx-xxx-xxx-no-such-host