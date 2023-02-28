# Unisearcher - Assignment-1
This project is a submission for the first assignment in PROG-2005: Cloud Technologies. 
Unisearcher is a web application written in Golang that utilizes REST API to provide information to 
the client about universities that may be good candidates for applications based on their name.
Additionally, it also provides convenient information related to the country the university is situated in.

## Endpoints
The web service has three resource root paths called: university_information (uniinfo), neighbour_universities (neighbourunis) and diag.

`/unisearcher/v1/uniinfo/`

`/unisearcher/v1/neighbourunis/`

`/unisearcher/v1/diag/`

### Uniinfo
The uniinfo endpoint returns information about the country/ies a university/ies containing a particular string in their name.

`Method: GET 
Path: /unisearcher/v1/uniinfo/
Request: uniinfo/{:partial_or_complete_university_name}{?{fields={:field1,field2,...}}&{limit={:number}}}`


#### Parameters
* `{:partial_or_complete_university_name}` partial or complete university name of the
universities to be searched for.

* `{?fields={:field1,field2,...}}` optional parameter. Specifies which fields to
be included in the result. Applicable fields are: name, country, isocode, webpages,
languages and map. The standard if there is no specified fields are all the fields.

* `{?limit={:number}}` optional parameter. Limits the number of universities displayed. The standard 
if there is no specified value is all results.

#### Response:
Content type: `application/json`

Status codes
* 200: Everything is OK.
* 204: No content.
* 400: Bad request.
* 404: Not found.
* 405: Method not allowed.
* 500: Internal server error.

#### Example requests and responses
Request: `uniinfo/stavanger`

`[{"name":"University of Stavanger","country":"Norway","isocode":"NO","webpages":["http://www.uis.no/"],"languages":{"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},"map":"https://www.openstreetmap.org/relation/2978650"}]`

Request: `uniinfo/stavanger?fields=country,isocode`

`[{"country":"Norway","isocode":"NO"}]`