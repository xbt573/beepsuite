# beepd
HTTP API for using motherboard speaker, currently only has stub (should have at least Linux in near future).

For now, it has only three routes:
* `/api/v1/beep`: Plain beep, nothing fancy.
* `/api/v1/beep/random`: Beeps with probability, given as `probability` query param.
* `/api/v1/beep/dice`: Beeps with 1/6 chance.

All routes can receive body with array of beeps to be made in format (example):
```jsonc
[
    {
        "frequency": 440,
        "length": 200
    },
    {
        "delay": 100
    }
    // ...
]
```

`frequency`/`length` and `delay` are mutually exclusive. If body is null, beepd performs single beep with 400Hz for 200ms.

For details, look at `requests.http` file in root directory of project.

beepd config is in this format (yaml):
```yaml
listen: 0.0.0.0:3000
```


This project does not have TLS and authentication by design, this will be handled externally in bounds of this project.