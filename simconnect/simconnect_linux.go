package simconnect

//go:generate go-bindata -pkg simconnect -o bindata.go -modtime 1 -prefix "../_vendor" "../_vendor/MSFS-SDK/SimConnect SDK/lib/SimConnect.dll"

// MSFS-SDK/SimConnect\ SDK/include/SimConnect.h
// MSFS-SDK/SimConnect\ SDK/lib/SimConnect.dll

// Everything is stubbed out on linux

import (
	"fmt"
	"reflect"
	"unsafe"
)

type SimConnect struct {
	handle      unsafe.Pointer
	DefineMap   map[string]DWORD
	LastEventID DWORD
}

func New(name string) (*SimConnect, error) {
	return nil, nil
}

func (s *SimConnect) GetEventID() DWORD {
	id := s.LastEventID
	s.LastEventID += 1
	return id
}

func (s *SimConnect) GetDefineID(a interface{}) DWORD {
	structName := reflect.TypeOf(a).Elem().Name()

	id, ok := s.DefineMap[structName]
	if !ok {
		id = s.DefineMap["_last"]
		s.DefineMap[structName] = id
		s.DefineMap["_last"] = id + 1
	}

	return id
}

func (s *SimConnect) RegisterDataDefinition(a interface{}) error {
	defineID := s.GetDefineID(a)

	v := reflect.ValueOf(a).Elem()
	for j := 1; j < v.NumField(); j++ {
		fieldName := v.Type().Field(j).Name
		nameTag, _ := v.Type().Field(j).Tag.Lookup("name")
		unitTag, _ := v.Type().Field(j).Tag.Lookup("unit")

		fieldType := v.Field(j).Kind().String()
		if fieldType == "array" {
			fieldType = fmt.Sprintf("[%d]byte", v.Field(j).Type().Len())
		}

		if nameTag == "" {
			return fmt.Errorf("%s name tag not found", fieldName)
		}

		dataType, err := derefDataType(fieldType)
		if err != nil {
			return err
		}

		s.AddToDataDefinition(defineID, nameTag, unitTag, dataType)
		//fmt.Printf("fieldName: %s  fieldType: %s  nameTag: %s unitTag: %s\n", fieldName, fieldType, nameTag, unitTag)
	}

	return nil
}

func (s *SimConnect) Close() error {
	// SimConnect_Open(
	//   HANDLE * phSimConnect,
	// );
	return nil
}

func (s *SimConnect) AddToDataDefinition(defineID DWORD, name, unit string, dataType DWORD) error {
	// SimConnect_AddToDataDefinition(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_DATA_DEFINITION_ID DefineID,
	//   const char * DatumName,
	//   const char * UnitsName,
	//   SIMCONNECT_DATATYPE DatumType = SIMCONNECT_DATATYPE_FLOAT64,
	//   float fEpsilon = 0,
	//   DWORD DatumID = SIMCONNECT_UNUSED
	// );
	return nil
}

func (s *SimConnect) SubscribeToSystemEvent(eventID DWORD, eventName string) error {
	// SimConnect_SubscribeToSystemEvent(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_CLIENT_EVENT_ID EventID,
	//   const char * SystemEventName
	// );

	return nil
}

func (s *SimConnect) RequestDataOnSimObjectType(requestID, defineID, radius, simobjectType DWORD) error {
	// SimConnect_RequestDataOnSimObjectType(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_DATA_REQUEST_ID RequestID,
	//   SIMCONNECT_DATA_DEFINITION_ID DefineID,
	//   DWORD dwRadiusMeters,
	//   SIMCONNECT_SIMOBJECT_TYPE type
	// );

	return nil
}

func (s *SimConnect) RequestDataOnSimObject(requestID, defineID, objectID, period, flags, origin, interval, limit DWORD) error {
	// SimConnect_RequestDataOnSimObject(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_DATA_REQUEST_ID RequestID,
	//   SIMCONNECT_DATA_DEFINITION_ID DefineID,
	//   SIMCONNECT_OBJECT_ID ObjectID,
	//   SIMCONNECT_PERIOD Period,
	//   SIMCONNECT_DATA_REQUEST_FLAG Flags = 0,
	//   DWORD origin = 0,
	//   DWORD interval = 0,
	//   DWORD limit = 0
	// );

	return nil
}

func (s *SimConnect) TransmitClientEvent(event, data DWORD) error {
	return nil
}

func (s *SimConnect) SetDataOnSimObject(defineID, simobjectType, flags, arrayCount, size DWORD, buf unsafe.Pointer) error {
	//s.SetDataOnSimObject(defineID, simconnect.OBJECT_ID_USER, 0, 0, size, buf)

	// SimConnect_SetDataOnSimObject(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_DATA_DEFINITION_ID DefineID,
	//   SIMCONNECT_OBJECT_ID ObjectID,
	//   SIMCONNECT_DATA_SET_FLAG Flags,
	//   DWORD ArrayCount,
	//   DWORD cbUnitSize,
	//   void * pDataSet
	// );

	return nil
}

func (s *SimConnect) SubscribeToFacilities(facilityType, requestID DWORD) error {
	// SimConnect_SubscribeToFacilities(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_FACILITY_LIST_TYPE type,
	//   SIMCONNECT_DATA_REQUEST_ID RequestID
	// );

	return nil
}

func (s *SimConnect) UnsubscribeToFacilities(facilityType DWORD) error {
	// SimConnect_UnsubscribeToFacilities(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_FACILITY_LIST_TYPE type
	// );

	return nil
}

func (s *SimConnect) RequestFacilitiesList(facilityType, requestID DWORD) error {
	// SimConnect_RequestFacilitiesList(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_FACILITY_LIST_TYPE type,
	//   SIMCONNECT_DATA_REQUEST_ID RequestID
	// );

	return nil
}

func (s *SimConnect) MapClientEventToSimEvent(eventID DWORD, eventName string) error {
	// SimConnect_MapClientEventToSimEvent(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_CLIENT_EVENT_ID EventID,
	//   const char * EventName = ""
	// );

	return nil
}

func (s *SimConnect) MenuAddItem(menuItem string, menuEventID, Data DWORD) error {
	// SimConnect_MenuAddItem(
	//   HANDLE hSimConnect,
	//   const char * szMenuItem,
	//   SIMCONNECT_CLIENT_EVENT_ID MenuEventID,
	//   DWORD dwData
	// );

	return nil
}

func (s *SimConnect) MenuDeleteItem(menuItem string, menuEventID, Data DWORD) error {
	// SimConnect_MenuDeleteItem(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_CLIENT_EVENT_ID MenuEventID
	// );

	return nil
}

func (s *SimConnect) AddClientEventToNotificationGroup(groupID, eventID DWORD) error {
	// SimConnect_AddClientEventToNotificationGroup(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
	//   SIMCONNECT_CLIENT_EVENT_ID EventID,
	//   BOOL bMaskable = FALSE
	// );

	return nil
}

func (s *SimConnect) SetNotificationGroupPriority(groupID, priority DWORD) error {
	// SimConnect_SetNotificationGroupPriority(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
	//   DWORD uPriority
	// );

	return nil
}

func (s *SimConnect) ShowText(textType DWORD, duration float64, eventID DWORD, text string) error {
	// SimConnect_Text(
	//   HANDLE hSimConnect,
	//   SIMCONNECT_TEXT_TYPE type,
	//   float fTimeSeconds,
	//   SIMCONNECT_CLIENT_EVENT_ID EventID,
	//   DWORD cbUnitSize,
	//   void * pDataSet
	// );

	return nil
}

func (s *SimConnect) GetNextDispatch() (unsafe.Pointer, int32, error) {

	return nil, 0, nil
}
