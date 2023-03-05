# Unisearcher - Assignment-1
This project is a submission for the first assignment in PROG-2005: Cloud Technologies.
Unisearcher is a web application written in Golang that utilizes REST API to provide information to
the client about universities that may be good candidates for applications based on their name.
Additionally, it also provides convenient information related to the country the university is situated in.

## Endpoints
The web service has three resource root paths called: university_information (uniinfo), neighbour_universities (neighbourunis) and diag.

```
/unisearcher/v1/uniinfo/
/unisearcher/v1/neighbourunis/
/unisearcher/v1/diag/
```

### Uniinfo
The uniinfo endpoint returns information about the country/ies a university/ies containing a particular string in their name.

```
Method: GET 
Path: /unisearcher/v1/uniinfo/
Request: uniinfo/{:partial_or_complete_university_name}{?{fields={:field1,field2,...}}&{limit={:number}}}
```

#### Parameters
* `{:partial_or_complete_university_name}` partial or complete university name of the
  universities to be searched for.

* `{?fields={:field1,field2,...}}` optional parameter. Specifies which fields to
  be included in the result. Applicable fields are: name, country, isocode, webpages,
  languages and map. The standard if there is no specified fields are all the fields.

* `{?limit={:number}}` optional parameter. Limits the number of universities displayed.
  If the amount of results are less than the limit this won't have an impact. The standard
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

Response:
```
[
    {
        "name": "University of Stavanger",
        "country": "Norway",
        "isocode": "NO",
        "webpages": [
            "http://www.uis.no/"
        ],
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        },
        "map": "https://www.openstreetmap.org/relation/2978650"
    }
]
```

Request: `uniinfo/stavanger?fields=country,isocode`

Response:
```
[
    {
        "country": "Norway",
        "isocode": "NO"
    }
]
```

Request: `uniinfo/stavanger?fields=languages&limit=1`

Response:
```
[
    {
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        }
    }
]
```

### Neighbourunis
The Neighbourunis endpoint provides an overview of universities in neighboring countries with the
same name component in their institution name.

```
Method: GET
Path: /unisearcher/v1/neighbourunis/
Request: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?{fields={:field1,field2,...}}&{limit={:number}}}
```

#### Parameters

* `{:country_name}` the country name in english that is the basis country of
  the search of universities with the identical name in neighbouring countries.

* `{:partial_or_complete_university_name}` partial or complete university name,
  for which universities with identical names in neighboring countries are sought.

* `{?fields={:field1,field2,...}}` optional parameter. Specifies which fields to
  be included in the result. Applicable fields are: name, country, isocode, webpages,
  languages and map. The standard if there is no specified fields are all the fields.

* `{?limit={:number}}` optional parameter. Limits the number of universities displayed.
  If the amount of results are less than the limit this won't have an impact. The standard
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

Request: `neighbourunis/norway/helsinki`

Response:
```
[
    {
        "name": "Helsinki School of Economics and Business Administration",
        "country": "Finland",
        "isocode": "FI",
        "webpages": [
            "http://www.hkkk.fi/"
        ],
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "map": "openstreetmap.org/relation/54224"
    },
    {
        "name": "Helsinki University of Technology",
        "country": "Finland",
        "isocode": "FI",
        "webpages": [
            "http://www.hut.fi/"
        ],
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "map": "openstreetmap.org/relation/54224"
    },
    {
        "name": "University of Art and Design Helsinki",
        "country": "Finland",
        "isocode": "FI",
        "webpages": [
            "http://www.uiah.fi/"
        ],
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "map": "openstreetmap.org/relation/54224"
    },
    {
        "name": "University of Helsinki",
        "country": "Finland",
        "isocode": "FI",
        "webpages": [
            "http://www.helsinki.fi/"
        ],
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "map": "openstreetmap.org/relation/54224"
    }
]
```

Request: `neighbourunis/sweden/science?fields=name`

Response:
```
[
    {
        "name": "Häme University of Applied Sciences"
    },
    {
        "name": "Laurea University of Applied Sciences"
    },
    {
        "name": "EVTEK University of Applied Sciences"
    },
    {
        "name": "Central Ostrobothnia University of Applied Sciences"
    },
    {
        "name": "Diaconia University of Applied Sciences"
    },
    {
        "name": "Rovaniemi University of Applied Sciences"
    },
    {
        "name": "Vaasa University of Applied Sciences"
    },
    {
        "name": "Metropolia University of Applied Sciences"
    }
]
```
Request: `neighbourunis/vietnam/technology?fields=name&limit=5`

Response:
```
[
    {
        "name": "University of Technology Phnom Penh"
    },
    {
        "name": "Chamreun University of Poly Technology"
    },
    {
        "name": "Nanjing University of Information Science and Technology"
    },
    {
        "name": "Changchun University of Science and Technology"
    },
    {
        "name": "Guangzhou College of South China University of Technology"
    }
]
```

### Diag
The diagnostics endpoint provides a simple health check. In other words the interface
indicates whether the individual services the service are dependent on are available or not.
The health check reports this information, as well as information about the uptime of the service.

```
Method: GET
Path: /unisearcher/v1/diag/
```

#### Response
Content type: `application/json`

Status codes
* 200: Everything is OK.
* 404: Not found.
* 500: Internal server error.

Diag content
* universitiesapi: the http status code for the "universities API".
* countriesapi: the http status code for "restcountries API".
* version: set to "v1".
* uptime: the time since the last service restart (in seconds).

#### Example request and response

Request: `/diag`

Response:

```
{
   "universitiesapi": "<http status code for universities API>",
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": <time elapsed in seconds from the last service restart>
}
```

Note: `<some value>` indicates placeholders for values to be populated by the service. 
An example response is provided underneath. 

Example response:
```
{
    "universitiesapi": "200",
    "countriesapi": "200",
    "version": "v1",
    "uptime": 5
}
```

## Deployment
The API is hosted on Render.

URL to the deployed Render service:
https://assignment-1-0tmo.onrender.com 

It can also be downloaded and run on local machine.

## Design

Throughout the implementation of this application, the focus points on the design has been loose coupling, high cohesion
and modularity as close to Golang convention as possible. This has been done through constants, different files for handlers and
generic functions.

### Project structure

The project structure was created with the goal of responsibility driven design,
and to minimize code duplication overall.

The endpoint-handlers got one file each, and are all located in the "handlers" package.
Since both uniinfo and neighbourunis need to request both countries and universities,
all code related to requesting and getting a response is moved to an own request package.

In order to limit API-requests, when countries are requested all borders from the country is retrieved and for each border request the country.
By doing this the API workload can get large. The API-server side's workload is reduced. In this way the REST-principles are met.

Extra features

* Cache: When searching for a country/ies, the country/ies will be cached for a set period of time. This is done to ensure an updated cache.
  This result in less frequent requests for the API and shorter response time.
* Query: The user is able to specify both the limit of results and which fields to be included in a response.

## Further improvements

These are further improvements I've not have had time to resolve.
* Use a middleware to set the content-type header for all response. 
* Implement Gorilla Mux to define URL routes and extract variables from them instead of doing it manually.
* When retrieving the neighbouring countries from a given country, a request is made for each bordering country. 
Instead, make a single request to retrieve all the countries that border the given country.
* Improve error handling: some errors happening in neighbourunis and uniinfo are not displaying. 
However, only showing "1", and a correct status code.
* Name differences - the two APIs has different writings for at least vietnam, decrease the lack of results. 
The list of countries this affects is not complete and the API would not give any warnings for this. 
Look for a solution to resolve this