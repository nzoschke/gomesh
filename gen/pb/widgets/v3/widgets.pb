
·
google/protobuf/empty.protogoogle.protobuf"
EmptyBv
com.google.protobufB
EmptyProtoPZ'github.com/golang/protobuf/ptypes/emptyø¢GPBªGoogle.Protobuf.WellKnownTypesbproto3
ê
 google/protobuf/field_mask.protogoogle.protobuf"!
	FieldMask
paths (	RpathsB‰
com.google.protobufBFieldMaskProtoPZ9google.golang.org/genproto/protobuf/field_mask;field_mask¢GPBªGoogle.Protobuf.WellKnownTypesbproto3
÷
google/protobuf/timestamp.protogoogle.protobuf";
	Timestamp
seconds (Rseconds
nanos (RnanosB~
com.google.protobufBTimestampProtoPZ+github.com/golang/protobuf/ptypes/timestampø¢GPBªGoogle.Protobuf.WellKnownTypesbproto3
¦
widgets/v3/widgets.protogomesh.widgets.v3google/protobuf/empty.proto google/protobuf/field_mask.protogoogle/protobuf/timestamp.proto"”
Widget
parent (	Rparent
name (	Rname!
display_name (	RdisplayName;
create_time (2.google.protobuf.TimestampR
createTime" 

GetRequest
name (	Rname"j
CreateRequest
parent (	Rparent
id (	Rid1
widget (2.gomesh.widgets.v3.WidgetRwidget"
UpdateRequest1
widget (2.gomesh.widgets.v3.WidgetRwidget;
update_mask (2.google.protobuf.FieldMaskR
updateMask"#
DeleteRequest
name (	Rname"a
ListRequest
parent (	Rparent
	page_size (RpageSize

page_token (	R	pageToken"k
ListResponse3
widgets (2.gomesh.widgets.v3.WidgetRwidgets&
next_page_token (	RnextPageToken"?
BatchGetRequest
parent (	Rparent
names (	Rnames"G
BatchGetResponse3
widgets (2.gomesh.widgets.v3.WidgetRwidgets2º
Widgets?
Get.gomesh.widgets.v3.GetRequest.gomesh.widgets.v3.WidgetE
Create .gomesh.widgets.v3.CreateRequest.gomesh.widgets.v3.WidgetE
Update .gomesh.widgets.v3.UpdateRequest.gomesh.widgets.v3.WidgetB
Delete .gomesh.widgets.v3.DeleteRequest.google.protobuf.EmptyG
List.gomesh.widgets.v3.ListRequest.gomesh.widgets.v3.ListResponseS
BatchGet".gomesh.widgets.v3.BatchGetRequest#.gomesh.widgets.v3.BatchGetResponseB-
com.gomesh.widgets.v3BWidgetsProtoPZv3pbbproto3