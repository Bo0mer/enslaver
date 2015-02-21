# enslaver
Master of all os-agents arount the world.

## Running it

### Configuration
Create configuration directory, which should contain `config.yml` file. You can preview example config file in the config/config.yml.

When you're done with the file, you should provide the config to the application by exporting a env variable `ENSLAVER_CONFIG_DIR` containing the full path to the directory, where your config.yml file is stored.

### Actual run
As simple as that:

```bash
go run main.go
```

## Running the tests
You need to have gingko installed.

```bash
ginkgo -r
```

## API

Following is the API that is provided by the OS-Agent for executing commands.

The response codes that are returned by the OS-Agent are splitted into the following groups:

| Group | Description |
| ---- | ----------- |
| 2XX | The requested operation completed successfully. |
| 4XX | There was a problem with the request payload. |
| 5XX | Execution of the operation failed due to some unexpected reason. |

**Payload**

The API is based on JSON request and responses. If not stated otherwise, default content-type should be `application/json`.

### Get Slaves

`GET /slave`

**Request**
Request should not contain any data. If it contains, it is discarded.

**Response**

| Name | Type | Description |
| ---- | ---- | ----------- |
| slaves | array | All registered slaves. |

Example Response:

```JSON
[
    {
        "id": "alice",
        "tags": {
            "tagKey1": "tagValue1",
            "tagKey2": "tagValue2"
        }
    },
    {
        "id": "bob",
        "tags": {}
    }
]
```

### Create Job

`POST /jobs`

**Request**

| Name | Type | Description |
| ---- | ---- | ----------- |
| slaves | array | Slaves which should execute the command. **Only id is required.** |
| async | boolean | This field is discarded for now. |
| command | struct | Properties of the command to run. |

Example Request:

```JSON
{
    "slaves": [
        {
            "id": "alice"
        },
        {
            "id": "bob"
        }
    ],
    "aysnc": true,
    "command": {
        "name": "cat",
        "args": [
            "arg1",
            "arg2"
        ],
        "env": {
            "variable_name1": "value1",
            "variable_name2": "value2"
        },
        "use_isolated_env": false,
        "working_dir": "/home/agent/whoa",
        "input": "This is the input to the cat command."
    }
}
```

**Response**

| Name | Type | Description |
| ---- | ---- | ----------- |
| id | string | The id of the (created) job. |
| status | string | The status of the job. Either `COMPLETED` or `IN_PROCESS`. |
| results | array | Contains an entry for each enslaver in the request. The entry specifies which is the slave, what is the status of it's execution, and the result of it. (see the example) |

**Note**
The job will be in status `COMPLETED` **iff** all the slaves have executed the command. 

Example Response:

```JSON
{
    "id": "job-id",
    "status": "COMPLETED",
    "results": [
        {
            "slave": {
                "id": "id-1",
                "tags": {
                    "tagKey": "tagValue",
                    "tagKey2": "tagValue2"
                }
            },
            "status": "COMPLETED",
            "result": {
                "stdout": "blabla",
                "stderr": "",
                "exit_code": 0,
                "error": ""
            }
        }
    ]
}
```

