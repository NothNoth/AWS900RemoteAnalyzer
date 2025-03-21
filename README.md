# AWS 900 Remote protocol analyzer

This is a Wireshark plugin for Solid State Logic remote, and more specifically the AWS900 traffic.

This Wireshark plugin uses [Wirego](https://github.com/quarkslab/wirego/).

## Getting started

Requirements:

  - Go framework
  - The wirego plugin and Wireshark properly set up (see Wirego's documentation)

Building:

  - make

Running:

  - run "./aws900"
  - run Wireshark
  - On the Wirego's preferences, make sur the endpoint **ipc:///tmp/wirego0** is set
  - restart Wireshark and open your pcap file

## State of work

Current version processes all packets seen during the discovery step (when the remote connects to the AWS console).

The reverse engineering is based on the AWS900Remote.jar application which is the source of this traffic.

