# secureapi

LoginUser

localhost:3453/login

[
 
{"UserName":"rezhnn",
    "Password":"123"   }

]

CreateNewAccount

localhost:3453/users/CreateUser

{

    "ResCode": "200",
    "ResMessage": "Success",
    "Data": {
        "NewUserId": "63a8102af373798c3a9daa7e"
    }
}


GetListAccounts

localhost:3453/users/getList

{

    "ResCode": "200",
    "ResMessage": "Success",
    "Data": [
        {
            "Id": "639ed31acaa94918231a89b5",
            "FirstName": "",
            "Code": "002528",
            "LastName": "",
            "BirthDate": "2002/10/3",
            "UserName": "rezhnn",
            "Password": "123"
        },
        {
            "Id": "639f0e863fc6cb51a15107f7",
            "FirstName": "reza",
            "Code": "002528",
            "LastName": "hayati",
            "BirthDate": "2002/10/13",
            "UserName": "rezhnn",
            "Password": "123"
        },
        {
            "Id": "63a8102af373798c3a9daa7e",
            "FirstName": "reza",
            "Code": "002528",
            "LastName": "hayati",
            "BirthDate": "2002/10/13",
            "UserName": "rezhnn",
            "Password": "123"
        }
    ]
}

