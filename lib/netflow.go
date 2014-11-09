//
// netflow.go
// Created by Daniel Piekacz on 2014-11-08.
// Updated on 2014-11-09.
// https://gixtools.net
//

package netflow

import (
//    "net"
)

const (
    NetFlowMessageID_TemplateV9 = 0
    NetFlowMessageID_TemplateV9_Optional = 1
    NetFlowMessageID_Template = 2
    NetFlowMessageID_Template_Optional = 3
    NetFlowMessageID_FlowRecord = 256
    NetFlowMessageID_Enterprise = 32768

    NetFlowTemplates_Version = 0
    NetFlowTemplates_Size = 1
    NetFlowTemplates_Template = 2
    NetFlowTemplates_Unpack = 3
    NetFlowTemplates_Struct = 4

    NetFlowDataTypes_In_Bytes = 1
    NetFlowDataTypes_In_Packets = 2
    NetFlowDataTypes_Flows = 3
    NetFlowDataTypes_Protocol = 4
    NetFlowDataTypes_Src_TOS = 5
    NetFlowDataTypes_TCP_Flags = 6
    NetFlowDataTypes_L4_Src_Port = 7
    NetFlowDataTypes_IPv4_Src_Addr = 8
    NetFlowDataTypes_Src_Mask = 9
    NetFlowDataTypes_Input_SNMP = 10
    NetFlowDataTypes_L4_Dst_Port = 11
    NetFlowDataTypes_IPv4_Dst_Addr = 12
    NetFlowDataTypes_Dst_Mask = 13
    NetFlowDataTypes_Output_SNMP = 14
    NetFlowDataTypes_IPv4_Next_Hop = 15
    NetFlowDataTypes_Src_AS = 16
    NetFlowDataTypes_Dst_AS = 17
    NetFlowDataTypes_BGP_IPv4_Next_Hop = 18
    NetFlowDataTypes_Mul_Dst_Packets = 19
    NetFlowDataTypes_Mul_Dst_Bytes = 20
    NetFlowDataTypes_Last_Switched = 21
    NetFlowDataTypes_First_Switched = 22
    NetFlowDataTypes_Out_Bytes = 23
    NetFlowDataTypes_Out_Packets = 24
    NetFlowDataTypes_Min_Packet_Length = 25
    NetFlowDataTypes_Max_Packet_Length = 26
    NetFlowDataTypes_IPv6_Src_Addr = 27
    NetFlowDataTypes_IPv6_Dst_Addr = 28
    NetFlowDataTypes_IPv6_Src_Mask = 29
    NetFlowDataTypes_IPv6_Dst_Mask = 30
    NetFlowDataTypes_IPv6_Flow_Label = 31
    NetFlowDataTypes_ICMP_Type = 32
    NetFlowDataTypes_Mul_IGMP_Type = 33
    NetFlowDataTypes_Sampling_Interval = 34
    NetFlowDataTypes_Sampling_Algorithm = 35
    NetFlowDataTypes_Flow_Active_Timeout = 36
    NetFlowDataTypes_Flow_Inactive_Timeout = 37
    NetFlowDataTypes_Engine_Type = 38
    NetFlowDataTypes_Engine_ID = 39
    NetFlowDataTypes_Total_Bytes_Exported = 40
    NetFlowDataTypes_Total_Packets_Exported = 41
    NetFlowDataTypes_Total_Flows_Exported = 42
//    NetFlowDataTypes_ = 43
    NetFlowDataTypes_IPv4_Src_Prefix = 44
    NetFlowDataTypes_IPv4_Dst_Prefix = 45
    NetFlowDataTypes_MPLS_Top_Label_Type = 46
    NetFlowDataTypes_MPLS_Top_Label_IP_Addr = 47
    NetFlowDataTypes_Flow_Sampler_ID = 48
    NetFlowDataTypes_Flow_Sampler_Mode = 49
    NetFlowDataTypes_Flow_Sampler_Random_Interval = 50
//    NetFlowDataTypes_ = 51
    NetFlowDataTypes_MinTTL = 52
    NetFlowDataTypes_MaxTTL = 53
    NetFlowDataTypes_IPv4_Ident = 54
    NetFlowDataTypes_Dst_TOS = 55
    NetFlowDataTypes_In_Src_MAC = 56
    NetFlowDataTypes_Out_Dst_MAC = 57
    NetFlowDataTypes_Src_VLAN = 58
    NetFlowDataTypes_Dst_VLAN = 59
    NetFlowDataTypes_IP_Proto_Version = 60
    NetFlowDataTypes_Direction = 61
    NetFlowDataTypes_IPv6_Next_Hop = 62
    NetFlowDataTypes_BGP_IPv6_Next_Hop = 63
    NetFlowDataTypes_IPv6_Option_Headers = 64
//    NetFlowDataTypes_ = 65
//    NetFlowDataTypes_ = 66
//    NetFlowDataTypes_ = 67
//    NetFlowDataTypes_ = 68
//    NetFlowDataTypes_ = 69
    NetFlowDataTypes_MPLS_Label1 = 70
    NetFlowDataTypes_MPLS_Label2 = 71
    NetFlowDataTypes_MPLS_Label3 = 72
    NetFlowDataTypes_MPLS_Label4 = 73
    NetFlowDataTypes_MPLS_Label5 = 74
    NetFlowDataTypes_MPLS_Label6 = 75
    NetFlowDataTypes_MPLS_Label7 = 76
    NetFlowDataTypes_MPLS_Label8 = 77
    NetFlowDataTypes_MPLS_Label9 = 78
    NetFlowDataTypes_MPLS_Label10 = 79
    NetFlowDataTypes_In_Dst_MAC = 80
    NetFlowDataTypes_Out_Src_MAc = 81
    NetFlowDataTypes_IF_Name = 82
    NetFlowDataTypes_IF_Desc = 83
    NetFlowDataTypes_Sampler_Name = 84
    NetFlowDataTypes_In_Permanent_Bytes = 85
    NetFlowDataTypes_In_Permanent_Packets = 86
//    NetFlowDataTypes_ = 87
    NetFlowDataTypes_Fragment_Offset = 88
    NetFlowDataTypes_Forwarding_status = 89
    NetFlowDataTypes_MPLS_PAL_Route_Distinguisher = 90
    NetFlowDataTypes_MPLS_Prefix_Length = 91
    NetFlowDataTypes_Src_Traffic_Index = 92
    NetFlowDataTypes_Dst_Traffic_Index = 93
    NetFlowDataTypes_Application_Desc = 94
    NetFlowDataTypes_Application_Tag = 95
    NetFlowDataTypes_Application_Name = 96
//    NetFlowDataTypes_ = 97
    NetFlowDataTypes_postipDiffServCodePoint = 98
    NetFlowDataTypes_Mul_Replication_Factor = 99
//    NetFlowDataTypes_ = 100
//    NetFlowDataTypes_ = 101
    NetFlowDataTypes_Layer2_Packet_Section_Offset = 102
    NetFlowDataTypes_Layer2_Packet_Section_Size = 103
    NetFlowDataTypes_Layer2_Packet_Section_Data = 104

    Protocols_ICMP = 1
    Protocols_TCP = 6
    Protocols_UDP = 17
    Protocols_IPV6 = 41
    Protocols_GRE = 47
    Protocols_ESP = 50
    Protocols_AH = 51
    Protocols_ICMP6 = 58
    Protocols_L2TP = 115
    Protocols_SCTP = 132

    TCPflags_FIN = 0x01
    TCPflags_SYN = 0x02
    TCPflags_RST = 0x04
    TCPflags_PSH = 0x08
    TCPflags_ACK = 0x10
    TCPflags_URG = 0x20
    TCPflags_ECE = 0x40
    TCPflags_CWR = 0x80

    // Never - For RFC special IP networks and known prefixes.
    PrefixExpire_Never = 0
    // 4 weeks - For prefixes where DNS lookup returned data.
    PrefixExpire_Default = 2419200
    // 2 hours - For prefixes where DNS lookup returned no data or failed.
    PrefixExpire_Short = 7200

    ASNtype_Internal = 0
    ASNtype_Unknown = 4294967295

    IP2ASN_def_mask_IPv4 = "24"
    IP2ASN_def_mask_IPv6 = "48"
)

type NetFlow_version struct {
    Version		uint16
}

type NetFlow_v1_header struct {
    Count		uint16
    System_Uptime	uint32
    Unix_seconds	uint32
    Unix_nseconds	uint32
}

type NetFlow_v1_data struct {
    Src_IP		[4]byte
    Dst_IP		[4]byte
    NextHop_IP		[4]byte
    In_Interface	uint16
    Out_Interface	uint16
    In_Packets		uint32
    In_Bytes		uint32
    Flow_First		uint32
    Flow_Last		uint32
    Src_Port		uint16
    Dst_Port		uint16
    Pad1		uint16
    Proto		uint8
    Src_TOS		uint8
    TCP_Flags		uint8
    Pad2		uint8
    Pad3		uint8
    Pad4		uint8
    Reserved		uint32
}

type NetFlow_v5_header struct {
    Count		uint16
    System_Uptime	uint32
    Unix_seconds	uint32
    Unix_nseconds	uint32
    Sequence_Number	uint32
    Engine_Type		uint8
    Engine_ID		uint8
    Sampling_Interval	uint16
}

type NetFlow_v5_data struct {
    Src_IP		[4]byte
    Dst_IP		[4]byte
    NextHop_IP		[4]byte
    In_Interface	uint16
    Out_Interface	uint16
    In_Packets		uint32
    In_Bytes		uint32
    Flow_First		uint32
    Flow_Last		uint32
    Src_Port		uint16
    Dst_Port		uint16
    Pad1		uint8
    TCP_Flags		uint8
    Proto		uint8
    Src_TOS		uint8
    Src_AS		uint16
    Dst_AS		uint16
    Src_Mask		uint8
    Dst_Mask		uint8
    Pad2		uint16
}

type NetFlow_v9_header struct {
    Count		uint16
    System_Uptime	uint32
    Unix_seconds	uint32
    Sequence_Number	uint32
    Source_ID		uint32
    Element_ID		uint16
    Field_Length	uint16
}

type NetFlow_v10_header struct {
    Message_Length	uint16
    Export_Timestamp	uint32
    Sequence_Number	uint32
    Domain_ID		uint32
    Element_ID		uint16
    Field_Length	uint16
}

type NetFlow_v10_enterprise struct {
    Enterprise_Number	uint32
}

type NetFlow_template_header struct {
    Template_ID		uint16
    Count		uint16
}

type NetFlow_template_data struct {
    Field_ID, Field_Size	uint16
}

type NetFlow_template struct {
    Count		uint16
    Format		[50]NetFlow_template_data
}
