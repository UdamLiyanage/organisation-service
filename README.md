# Organisation Microservice
Build Status: ![](https://github.com/UdamLiyanage/organisation-service/workflows/Go/badge.svg)
***
Organisation Microservice for IoT Platform. This service is responsible to the following features:
*Create Organisations
*Read Organisation data
*Update Organisation data
*Delete Organisations

***
## Document Structure for User - Revision 1
Below is the structure of the JSON document that holds organisation data
```
{
	"name": "Organisation Name",
	"devices": [{
			"device_name": "Device Name",
			"device_id": "Device ID"
		},
		{
			"device_name": "Device Name",
			"device_id": "Device ID"
		},
		{
			"device_name": "Device Name",
			"device_id": "Device ID"
		}
	],
	"users": [{
			"user_name": "User Name",
			"user_id": "User ID"
		},
		{
			"user_name": "User Name",
			"user_id": "User ID"
		},
		{
			"user_name": "User Name",
			"user_id": "User ID"
		}
	],
	"created_at": "Created Time",
	"updated_at": "Updated Time"
}
```
