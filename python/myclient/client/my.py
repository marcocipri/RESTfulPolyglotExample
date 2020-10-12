from __future__ import print_function

import time
import RESTfulPolyglotEamplePythonClient
from RESTfulPolyglotEamplePythonClient.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://petstore.swagger.io/api
# See configuration.py for a list of all supported configuration parameters.
configuration = RESTfulPolyglotEamplePythonClient.Configuration(
    host = "http://localhost:8080/api"
)



# Enter a context with an instance of the API client
with RESTfulPolyglotEamplePythonClient.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = RESTfulPolyglotEamplePythonClient.DefaultApi(api_client)

    try:
        api_response = api_instance.find_pets(limit=2)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DefaultApi->find_pet: %s\n" % e)
    