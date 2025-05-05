# Lab \- Use RESTCONF to Access a Networking Device

## Objectives

**Part 1: Install and Configure Docker Desktop and Postman**

**Part 2: Use Postman to Send GET Requests**

**Part 3: Use Postman to Send a PUT Request**

**Part 4: Use a Python Script to Send GET Requests**

**Part 5: Use a Python Script to Send a PUT Request**

## Background / Scenario

The RESTCONF protocol provides a simplified subset of NETCONF features over a RESTful API.\
RESTCONF allows you to make RESTful API calls to an networking device. The data returned by\
the API can be formatted in either XML or JSON. In the first half of this lab, you will use \
the Postman program to construct and send API requests to the RESTCONF service that is running\
on the Networking device emulator. In the second half of the lab, you will create Python scripts\
to perform the same tasks as your Postman program.

### Required Resources

* PC with operating system of your choice

##    1. Install Docker Desktop and start the RESTCONF container

In this Part, you install the Docker Desktop and start RESTCONF container.\
Also you will install and configure the Postman.

If you have not already completed the\
**Lab \- Install the Docker Descktop**, do so now. \
If you have already completed that lab, start the [Docker container for RESTCONF](TBD) now.

In the next part, you will install and open [Postman](https://www.postman.com/),\
disable SSL certificates, and explore the user interface.

*   Install the the Postman application following the [instruction](https://www.postman.com/downloads/)
* 	If this is the first time you have opened Postman, it may ask you to create an account or sign in. At the bottom of the window, you can also click the “Skip” message to skip signing in. Signing in is not required to use this application.
* 	By default, Postman has SSL certification verification turned on. You will not be using SSL certificates with the Restconf application; therefore, you need to turn off this feature.
     - Click **File \> Settings**.
     - Under the **General** tab, set the **SSL certificate verification** to **OFF**.
     - Close the **Settings** dialog box.

##  2. Use Postman to Send GET Requests

In this Part, you will use Postman to send a GET request to the \
RESTCONF container emulating networking device to verify that you can connect to the RESTCONF service.

**Explore the Postman user interface.**

In the center, you will see the **Launchpad**. You can explore this area if you wish.
Click the plus sign (+) next to the **Launchpad** tab to open a **GET Untitled Request**. This interface is where you will do all of your work in this lab.

* Enter the URL for the RESTCONF container

* The request type is already set to GET. Leave the request type set to GET.

* In the “Enter request URL” field, type in the URL that will be used to access the RESTCONF service that is running on container.

```
   https://localhost:8080/restconf
```

* Enter authentication credentials.

Under the URL field, there are tabs listed for **Params**, **Authorization**, **Headers**, **Body**, **Pre-request Script**, **Test**, and **Settings**. In this lab, you will use **Authorization**, **Headers**, and **Body**.

* Click the **Authorization** tab.
    * Under Type, click the down arrow next to “Inherit auth from parent” and choose **Basic Auth**.
    * For **Username** and **Password**, enter the local authentication credentials for the RESTCONF service that is running on container:

```
Username: admin / Password: password123
```

* Click **Headers**. Then click the **7 hidden**. You can verify that the Authorization key has a Basic value that will be used to authenticate the request when it is sent to the RESTCONF service.

* Set JSON as the data type to send to and receive from the RESTCONF service.

You can send and receive data from the RESTCONF service in XML or JSON format. For this lab, you will use JSON.

- In the **Headers** area, click in the first blank **Key** field and type **Content-Type** for the type of key. In the **Value** field, type **application/yang-data+json**. This tells Postman to send JSON data to the RESTCONF service.

- Below your **Content-Type** key, add another key/value pair. The **Key** field is **Accept** and the **Value** field is **application/yang-data+json**.

**Note**: You can change application/yang-data+json to application/yang-data+xml to send and receive XML data instead of JSON data, if necessary.

* Send the API request to the RESTCONF service.

Postman now has all the information it needs to send the GET request. Click **Send**. Below **Temporary Headers**, you should see the following JSON response from the RESTCONF service. If not, verify that you completed the previous steps in this part of the lab and correctly configured RESTCONF and HTTPS service in Part 2\.

```
{
    "ietf-restconf:restconf": " ",
    "data": "{}",
    "operations": "{}",
    "yang-library-version": "2016-06-21"
}
```

This JSON response verifies that Postman can now send other REST API requests to the RESTCONF service emulating network device.

**Use a GET request to gather the information for all interfaces on the network device.**

Now that you have a successful GET request, you can use it as a template for additional requests. At the top of Postman, next to the **Launchpad** tab, right-click the **GET** tab that you just used and choose **Duplicate Tab**.
Use the **ietf-interfaces** YANG model to gather interface information. For the URL, add **data/ietf-interfaces:interfaces**:

```
  http://localhost:8080/restconf/data/ietf-interfaces:interfaces
```

*   Click **Send**. You should see a JSON response from the RESTCONF service that is similar to the output shown below. Your output may be different depending on your particular router.

```
[
    {
        "ietf-interfaces:interfaces": {
            "interface": [
                {
                    "name": "GigabitEthernet1",
                    "description": "Main Interface",
                    "type": "iana-if-type:ethernetCsmacd",
                    "enabled": true,
                    "ietf-ip:ipv4": {},
                    "ietf-ip:ipv6": {}
                },
                {
                    "name": "Loopback1",
                    "description": "Loopback Interface",
                    "type": "iana-if-type:softwareLoopback",
                    "enabled": true,
                    "ietf-ip:ipv4": {},
                    "ietf-ip:ipv6": {}
                }
            ]
        }
    }
]
```
**Use a GET request to gather information for a specific interface on the network device.**

In this lab, only the GigabitEthernet1 interface is configured. To specify just this interface,\
extend the URL to only request information for this interface.

*   Duplicate your last GET request. Add the **interface=** parameter to specify an interface and type in the name of the interface.

```
http://localhost:8080/restconf/data/ietf-interfaces:interfaces/interface=GigabitEthernet1
```

*   Click **Send**. You should see a JSON response from the network devices that is similar to output below.\
Your output may be different depending on your particular router.

```
{
    "name": "GigabitEthernet1",
    "description": "Main Interface",
    "type": "iana-if-type:ethernetCsmacd",
    "enabled": true,
    "ietf-ip:ipv4": {},
    "ietf-ip:ipv6": {}
}
```

## **Use Postman to Send a PUT Request**

In this Part, you will configure Postman to send a PUT request to the Cat8000v to create a new loopback interface.

**Note**: If you created a Loopback interface in another lab, either remove it now or create a new one by using a different number.

1. ### **Duplicate and modify the last GET request.**

    1. Duplicate the last GET request.

        2. For the **Type** of request, click the down arrow next to **GET** and choose **PUT**.

            3. For the **interface=** parameter, change it to **\=Loopback1** to specify a new interface.

   [https://10.10.20.48/restconf/data/ietf-interfaces:interfaces/interface=Loopback1](https://10.10.20.48/restconf/data/ietf-interfaces:interfaces/interface=Loopback1)

    2. ### **Configure the body of the request specifying the information for the new loopback.**

        1. To send a PUT request, you need to provide the information for the body of the request. Next to the **Headers** tab, click **Body**. Then click the **Raw** radio button. The field is currently empty. If you click **Send** now, you will get error code **400 Bad Request** because Loopback1 does not exist yet and you did not provide enough information to create the interface.

        2. Fill in the **Body** section with the required JSON data to create a new Loopback1 interface. You can copy the Body section of the previous GET request and modify it. Or you can copy the following into the Body section of your PUT request. Notice that the type of interface must be set to **softwareLoopback**.

   {

   "ietf-interfaces:interface": {

       "name": "Loopback1",

       "description": "My first RESTCONF loopback",

       "type": "iana-if-type:softwareLoopback",

       "enabled": true,

       "ietf-ip:ipv4": {

         "address": \[

           {

             "ip": "10.1.1.1",

             "netmask": "255.255.255.0"

           }

         \]

       },

       "ietf-ip:ipv6": {}

   }

   }

         3. Click **Send** to send the PUT request to the Cat8000v. Below the Body section, you should see the HTTP response code **Status:** **201 Created**. This indicates that the resource was created successfully.

         4. You can verify that the interface was created. Return to your SSH session with the Cat8000v and enter **show ip interface brief**. You can also run the Postman tab that contains the request to get information about the interfaces on the Cat8000v that was created in the previous Part of this lab.

*Open configuration window*

Cat8000v\# **show ip interface brief**

Interface              IP-Address      OK? Method Status                Protocol

GigabitEthernet1       10.10.20.48     YES manual up                    up

Loopback1              10.1.1.1        YES other  up                    up

Cat8000v\#

*Close configuration window*

6. ## **Use a Python script to Send GET Requests**

In this Part, you will create a Python script to send GET requests to the Cat8000v.

1. ### **Create the RESTCONF directory and start the script.**

    1. Open VS code. Then click **File \> Open** **Folder...** and navigate to the **devnet-src** directory. Click **OK**.

        2. Open a terminal window in VS Code: **Terminal \> New Terminal**.

            3. Create a subdirectory called **restconf** in the **/devnet-src** directory.

   devasc@labvm:\~/labs/devnet-src$ **mkdir restconf**

   devasc@labvm:\~/labs/devnet-src$

         4. In the **EXPLORER** pane under **DEVNET-SRC**, right-click the **restconf** directory and choose **New File**.

         5. Name the file **restconf-get.py**. 

         1. Enter the following commands to import the modules that are required and disable SSL certificate warnings:

   import json

   import requests

   requests.packages.urllib3.disable\_warnings()

   The **json** module includes methods to convert JSON data to Python objects and vice versa. The **requests** module has methods that will let you send REST requests to a URL.

    2. ### **Create the variables that will be the components of the request.**

        1. Create a variable named **api\_url** and assign it the URL that will access the interface information on the Cat8000v :

   api\_url \= "https://10.10.20.48/restconf/data/ietf-interfaces:interfaces"



         2. Create a dictionary variable named **headers** that has keys for **Accept** and **Content-type** and assign the keys the value **application/yang-data+json**.

headers \= { "Accept": "application/yang-data+json",

               "Content-type":"application/yang-data+json"

              }

         3. Create a Python tuple variable named **basicauth** that has two keys needed for authentication, **username** and **password**.

basicauth \= ("developer", "C1sco12345")



      3. ### **Create a variable to send the request and store the JSON response.**

Use the variables that were created in the previous step as parameters for the **requests.get()** method. This method sends an HTTP GET request to the RESTCONF API on the Cat8000v. Assign the result of the request to a variable named **resp**. That variable will hold the JSON response from the API. If the request is successful, the JSON will contain the returned YANG data model.

1. Enter the following statement:

   resp \= requests.get(api\_url, auth=basicauth, headers=headers, verify=False)

The table below lists the various elements of this statement:

| Element | Explanation |
| :---- | :---- |
| resp | The variable to hold the response from the API |
| requests.get() | The method that actually makes the GET request |
| api\_url | The variable that holds the URL address string |
| auth | The tuple variable created to hold the authentication information |
| headers=headers | A parameter that is assigned the headers variable |
| verify=False | Disables verification of the SSL certificate when the request is made |

2. To see the HTTP response code, add a print statement.

   print(resp)

    3. Save and run your script. You should get the output shown below. If not, verify all previous steps in this part as well as the SSH and RESTCONF configuration for the Cat8000v.

   devasc@labvm:\~/labs/devnet-src$ **cd restconf/**

   devasc@labvm:\~/labs/devnet-src/restconf$ **python3 restconf-get.py**

   \<Response \[200\]\>

   devasc@labvm:\~/labs/devnet-src/restconf$

    4. ### **Format and display the JSON data received from the Cat8000v.**

Now you can extract the YANG model response values from the response JSON.

1. The response JSON is not compatible with Python dictionary and list objects, so it must be converted to Python format. Create a new variable called **response\_json** and assign the variable **resp** to it. Add the **json()** method to convert the JSON. The statement is as follows:

   response\_json \= resp.json()

    2. Add a print statement to display the JSON data.

   print(response\_json)

    3. Save and run your script. You should get output similar to the following:

   devasc@labvm:\~/labs/devnet-src/restconf$ **python3 restconf-get.py**

   \<Response \[200\]\>

   {'ietf-interfaces:interfaces': {'interface': \[{'name': 'GigabitEthernet1', 'description': 'VBox', 'type': 'iana-if-type:ethernetCsmacd', 'enabled': True, 'ietf-ip:ipv4': {'address': \[{'ip': '10.10.20.48', 'netmask': '255.255.255.0'}\]}, 'ietf-ip:ipv6': {}}, {'name': 'Loopback1', 'description': 'My first RESTCONF loopback', 'type': 'iana-if-type:softwareLoopback', 'enabled': True, 'ietf-ip:ipv4': {'address': \[{'ip': '10.1.1.1', 'netmask': '255.255.255.0'}\]}, 'ietf-ip:ipv6': {}}\]}}

   devasc@labvm:\~/labs/devnet-src/restconf$

         4. To prettify the output, edit your print statement to use the **json.dumps()** function with the “indent” parameter:

   print(json.dumps(response\_json, indent=4))

         5. Save and run your script. You should get the output shown below. This output is virtually identical to the output of your first Postman GET request.

   devasc@labvm:\~/labs/devnet-src/restconf$ **python3 restconf-get.py**

   \<Response \[200\]\>

   {

       "ietf-interfaces:interfaces": {

           "interface": \[

               {

                   "name": "GigabitEthernet1",

                   "description": "VBox",

                   "type": "iana-if-type:ethernetCsmacd",

                   "enabled": true,

                   "ietf-ip:ipv4": {

                       "address": \[

                           {

                               "ip": "10.10.20.48",

                               "netmask": "255.255.255.0"

                           }

                       \]

                   },

                   "ietf-ip:ipv6": {}

               },

               {

                   "name": "Loopback1",

                   "description": "My first RESTCONF loopback",

                   "type": "iana-if-type:softwareLoopback",

                   "enabled": true,

                   "ietf-ip:ipv4": {

                       "address": \[

                           {

                               "ip": "10.1.1.1",

                               "netmask": "255.255.255.0"

                           }

                       \]

                   },

                   "ietf-ip:ipv6": {}

               }

           \]

       }

   }

   devasc@labvm:\~/labs/devnet-src/restconf$

    7. ## **Use a Python Script to Send a PUT Request**

In this Part, you will create a Python script to send a PUT request to the Cat8000v. As was done in Postman, you will create a new loopback interface.

1. ### **Import modules and disable SSL warnings.**

    1. In the **EXPLORER** pane under **DEVNET-SRC**, right-click the **restconf** directory and choose **New File**.

        2. Name the file **restconf-put.py**.

            3. Enter the following commands to import the modules that are required and disable SSL certificate warnings:

   import json

   import requests

   requests.packages.urllib3.disable\_warnings()

    2. ### **Create the variables that will be the components of the request.**

        1. Create a variable named **api\_url** and assign it the URL that targets a new Loopback2 interface.

   **Note**: This variable specification should be on one line in your script.

   api\_url \= "https://10.10.20.48/restconf/data/ietf-interfaces:interfaces/interface=Loopback2"

         2. Create a dictionary variable named **headers** that has keys for Accept and Content-type and assign the keys the value **application/yang-data+json**.

   headers \= { "Accept": "application/yang-data+json",

               "Content-type":"application/yang-data+json"

              }

         3. Create a Python tuple variable named **basicauth** that has two values needed for authentication, username and password.

   basicauth \= ("developer", "C1sco12345")



         4. Create a Python dictionary variable **yangConfig** that will hold the YANG data that is required to create the new interface Loopback2. You can use the same dictionary that you used previously in Postman. However, change the interface number and address. Also, be aware that Boolean values must be capitalized in Python. Therefore, make sure that the **T** is capitalized in the key/value pair for **“enabled”: True**.

yangConfig \= {

       "ietf-interfaces:interface": {

           "name": "Loopback2",

           "description": "My second RESTCONF loopback",

           "type": "iana-if-type:softwareLoopback",

           "enabled": True,

           "ietf-ip:ipv4": {

               "address": \[

                   {

                       "ip": "10.2.1.1",

                       "netmask": "255.255.255.0"

                   }

               \]

           },

           "ietf-ip:ipv6": {}

       }

}

      3. ### **Create a variable to send the request and store the JSON response.**

Use the variables created in the previous step as parameters for the **requests.put()** method. This method sends an HTTP PUT request to the RESTCONF API. Assign the result of the request to a variable named **resp**. That variable will hold the JSON response from the API. If the request is successful, the JSON will contain the returned YANG data model.

1. Before entering statements, please note that this variable specification should be on only one line in your script. Enter the following statements:

   **Note**: This variable specification should be on one line in your script.

   resp \= requests.put(api\_url, data=json.dumps(yangConfig), auth=basicauth, headers=headers, verify=False)

    2. Enter the code below to handle the response. If the response is one of the HTTP success messages, the first message will be printed. Any other code value is considered an error. The response code and error message will be printed in the event that an error has been detected.

   if(resp.status\_code \>= 200 and resp.status\_code \<= 299):

       print("STATUS OK: {}".format(resp.status\_code))

   else:

       print('Error. Status Code: {} \\nError message: {}'.format(resp.status\_code,resp.json()))

The table below lists the various elements of these statements:

| Element | Explanation |
| :---- | :---- |
| resp | The variable to hold the response from the API. |
| requests.put() | The method that makes the PUT request. |
| api\_url | The variable that holds the URL address string. |
| data | The data to be sent to the API endpoint, which is formatted as JSON. |
| auth | The tuple variable created to hold the authentication information. |
| headers=headers | A parameter that is assigned the headers variable. |
| verify=False | A parameter that disables verification of the SSL certificate when the request is made. |
| resp.status\_code | The HTTP status code in the API **PUT** request reply. |

3. Save and run the script to send the PUT request to the Cat8000v. You should get a **201 Status Created** message. If not, check your code and the configuration for the Cat8000v.

    4. You can verify that the interface was created by entering **show ip interface brief** on the Cat8000v.

*Open configuration window*

Cat8000v\# **show ip interface brief**

Interface              IP-Address      OK? Method Status                Protocol

GigabitEthernet1       10.10.20.48     YES manual up                    up

Loopback1              10.1.1.1        YES other  up                    up

Loopback2              10.2.1.1        YES other  up                    up

Cat8000v\#

*Close configuration window*

3. # **Programs Used in this Lab**

The following Python scripts were used in this lab:

\#===================================================================

\#resconf-get.py

import json

import requests

requests.packages.urllib3.disable\_warnings()

api\_url \= "https://10.10.20.48/restconf/data/ietf-interfaces:interfaces"

headers \= { "Accept": "application/yang-data+json",

            "Content-type":"application/yang-data+json"

           }

basicauth \= ("cisco", "C1sco12345")

resp \= requests.get(api\_url, auth=basicauth, headers=headers, verify=False)

print(resp)

response\_json \= resp.json()

print(json.dumps(response\_json, indent=4))

\#end of file

\#===================================================================

\#resconf-put.py

import json

import requests

requests.packages.urllib3.disable\_warnings()

api\_url \= "https://10.10.20.48/restconf/data/ietf-interfaces:interfaces/interface=Loopback2"

headers \= { "Accept": "application/yang-data+json",

            "Content-type":"application/yang-data+json"

           }

basicauth \= ("cisco", "C1sco12345")

yangConfig \= {

    "ietf-interfaces:interface": {

        "name": "Loopback2",

        "description": "My second RESTCONF loopback",

        "type": "iana-if-type:softwareLoopback",

        "enabled": True,

        "ietf-ip:ipv4": {

            "address": \[

                {

                    "ip": "10.2.1.1",

                    "netmask": "255.255.255.0"

                }

            \]

        },

        "ietf-ip:ipv6": {}

    }

}

resp \= requests.put(api\_url, data=json.dumps(yangConfig), auth=basicauth, headers=headers, verify=False)

if(resp.status\_code \>= 200 and resp.status\_code \<= 299):

    print("STATUS OK: {}".format(resp.status\_code))

else:

    print('Error. Status Code: {} \\nError message: {}'.format(resp.status\_code,resp.json()))

\#end of file

*End of document*