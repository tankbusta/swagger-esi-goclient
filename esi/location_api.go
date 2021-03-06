/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.1.dev2
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"net/url"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type LocationApiService service


/**
 * Get character location
 * Information about the characters current location. Returns the current solar system id, and also the current station or structure ID if applicable.  ---  Alternate route: &#x60;/v1/characters/{character_id}/location/&#x60;  Alternate route: &#x60;/legacy/characters/{character_id}/location/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/location/&#x60;   ---  This route is cached for up to 5 seconds
 *
 * @param characterId An EVE character ID 
 * @param datasource(nil) The server name you would like data from 
 * @return *GetCharactersCharacterIdLocationOk
 */
func (a LocationApiService) GetCharactersCharacterIdLocation(characterId int32, datasource interface{}) (*GetCharactersCharacterIdLocationOk,  error) {
	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/location/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(evesso)' required
	// oauth required
	if a.client.Config.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.client.Config.AccessToken
	}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.client.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	 var successPayload = new(GetCharactersCharacterIdLocationOk)

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarHttpContentType)
	 if err != nil {
		  return successPayload, err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, err
	 }

	 defer localVarHttpResponse.Body.Close()
	 buf, err := ioutil.ReadAll(localVarHttpResponse.Body)
	 if err != nil {
		  return nil, err
	 }

	 err = json.Unmarshal([]byte(buf), &successPayload)

	return successPayload, err
}

/**
 * Get current ship
 * Get the current ship type, name and id  ---  Alternate route: &#x60;/v1/characters/{character_id}/ship/&#x60;  Alternate route: &#x60;/legacy/characters/{character_id}/ship/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/ship/&#x60;   ---  This route is cached for up to 5 seconds
 *
 * @param characterId An EVE character ID 
 * @param datasource(nil) The server name you would like data from 
 * @return *GetCharactersCharacterIdShipOk
 */
func (a LocationApiService) GetCharactersCharacterIdShip(characterId int32, datasource interface{}) (*GetCharactersCharacterIdShipOk,  error) {
	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/ship/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(evesso)' required
	// oauth required
	if a.client.Config.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.client.Config.AccessToken
	}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.client.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	 var successPayload = new(GetCharactersCharacterIdShipOk)

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarHttpContentType)
	 if err != nil {
		  return successPayload, err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, err
	 }

	 defer localVarHttpResponse.Body.Close()
	 buf, err := ioutil.ReadAll(localVarHttpResponse.Body)
	 if err != nil {
		  return nil, err
	 }

	 err = json.Unmarshal([]byte(buf), &successPayload)

	return successPayload, err
}

