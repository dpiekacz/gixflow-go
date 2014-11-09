//
// gixglow.go
// Created by Daniel Piekacz on 2014-11-08.
// Updated on 2014-11-09.
// https://gixtools.net
//
package main

import (
    "fmt"
    "net"
    "bytes"
    "log"
    "runtime"
    "time"
    "encoding/binary"
    "./lib"
)

func handlePacket(i int, pkt_src_ip net.IP, data []byte, pkt_size int, nf_templates map[string]netflow.NetFlow_template) {
    var nf_version netflow.NetFlow_version

    var nf_v1_head netflow.NetFlow_v1_header
    var nf_v1_data netflow.NetFlow_v1_data

    var nf_v5_head netflow.NetFlow_v5_header
    var nf_v5_data netflow.NetFlow_v5_data

    var nf_v9_head netflow.NetFlow_v9_header

    var nf_v10_head netflow.NetFlow_v10_header
    var nf_v10_ent  netflow.NetFlow_v10_enterprise

    var nf_tmpl_head netflow.NetFlow_template_header
    var nf_tmpl_data netflow.NetFlow_template_data


    buf := bytes.NewReader(data)
    binary.Read(buf, binary.BigEndian, &nf_version)

    switch nf_version.Version {
    case 1:
	binary.Read(buf, binary.BigEndian, &nf_v1_head)
	for i := 0; i < int(nf_v1_head.Count); i++ {
		binary.Read(buf, binary.BigEndian, &nf_v1_data)
//		fmt.Println(i, runtime.NumGoroutine(), pkt_src_ip, nf_version.Version, pkt_size, nf_v1_head, nf_v1_data)
	}

    case 5:
	binary.Read(buf, binary.BigEndian, &nf_v5_head)
	for i := 0; i < int(nf_v5_head.Count); i++ {
		binary.Read(buf, binary.BigEndian, &nf_v5_data)
//		fmt.Println(i, runtime.NumGoroutine(), pkt_src_ip, nf_version.Version, pkt_size, nf_v5_head, nf_v5_data)
	}

    case 9:
	binary.Read(buf, binary.BigEndian, &nf_v9_head)

	if nf_v9_head.Element_ID == netflow.NetFlowMessageID_TemplateV9 {
	    for i := 0; i < int(nf_v9_head.Count); i++ {
		var nf_tmpl netflow.NetFlow_template

		binary.Read(buf, binary.BigEndian, &nf_tmpl_head)
		nf_tmpl.Count = nf_tmpl_head.Count

		for j := 0; j < int(nf_tmpl_head.Count); j++ {
		    binary.Read(buf, binary.BigEndian, &nf_tmpl_data)
		    nf_tmpl.Format[j] = nf_tmpl_data
		}

		nf_index := pkt_src_ip.String() + "-v9-" + fmt.Sprintf("%v", nf_tmpl_head.Template_ID) + "-0"
		nf_templates[nf_index] = nf_tmpl
//		fmt.Println(nf_templates)
	    }
	}

	if nf_v9_head.Element_ID >= netflow.NetFlowMessageID_FlowRecord {
	    nf_index := pkt_src_ip.String() + "-v9-" + fmt.Sprintf("%v", nf_v9_head.Element_ID) + "-0"
	    nf_tmpl, ok := nf_templates[nf_index]
	    if ok {
		fmt.Println("data v9 - template found", nf_index, nf_tmpl)
	    } else {
		fmt.Println("data v9 - no template", nf_index)
	    }
	}

    case 10:
	binary.Read(buf, binary.BigEndian, &nf_v10_head)

	if nf_v10_head.Element_ID & netflow.NetFlowMessageID_Enterprise == netflow.NetFlowMessageID_Enterprise {
	    binary.Read(buf, binary.BigEndian, &nf_v10_ent)
	}

	if nf_v10_head.Element_ID == netflow.NetFlowMessageID_Template {
	    //for i := 0; i < int(nf_v10_head.Field_Length); i++ {
	    var nf_tmpl netflow.NetFlow_template

	    binary.Read(buf, binary.BigEndian, &nf_tmpl_head)
	    nf_tmpl.Count = nf_tmpl_head.Count

	    for j := 0; j < int(nf_tmpl_head.Count); j++ {
	        binary.Read(buf, binary.BigEndian, &nf_tmpl_data)
	        nf_tmpl.Format[j] = nf_tmpl_data
	    }

	    nf_index := pkt_src_ip.String() + "-v10-" + fmt.Sprintf("%v", nf_tmpl_head.Template_ID) + "-" + fmt.Sprintf("%v", nf_v10_head.Domain_ID)
	    nf_templates[nf_index] = nf_tmpl
	}

	if nf_v10_head.Element_ID >= netflow.NetFlowMessageID_FlowRecord {
	    nf_index := pkt_src_ip.String() + "-v10-" + fmt.Sprintf("%v", nf_v10_head.Element_ID) + "-" + fmt.Sprintf("%v", nf_v10_head.Domain_ID)
	    nf_tmpl, ok := nf_templates[nf_index]
	    if ok {
		fmt.Println("data v10 - template found", nf_index, nf_tmpl)
	    } else {
		fmt.Println("data v10 - no template", nf_index)
	    }
	}

    default:
	fmt.Println(i, runtime.NumGoroutine(), pkt_src_ip, nf_version.Version, pkt_size, "Unknown Netflow version")
    }
}

func netflow_receiver() {
    udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:19000")
    if err != nil {
	log.Fatalf("ResolveUDPAddr failed: %s\n", err)
    }

    sock, err := net.ListenUDP("udp", udpAddr)
    if err != nil {
	log.Fatalf("ListenUDP failed: %s\n", err.Error())
    }

    nf_templates := make(map[string]netflow.NetFlow_template)

    i := 0
    for {
	i++
        pkt_buf := make([]byte, 1500)
	pkt_size, pkt_src_ip, err := sock.ReadFromUDP(pkt_buf)
	if err != nil {
		fmt.Println(err)
	}
        go handlePacket(i, pkt_src_ip.IP, pkt_buf, pkt_size, nf_templates)
    }
}


func main() {
    netflow_receiver()
    for {
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
    }
}
