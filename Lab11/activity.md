# Laboratório de Redes - Aula 11

## Capturing packets from an execution of traceroute && A look at the ICMP protocol

1. Select the first ICMP Echo Request message sent by your computer, and expand the Internet Protocol part of the packet in the pacet details window. What is the IP address of your computer?

    **R:** The IP is 192.168.0.111

2. Within the IP packet header, what is the value in the upper layer protocol field?

    **R:** The value is 1 (ICMP) which is the protocol number for ICMP.

3. How many bytes are in the IP header? How many bytes are in the payload of the IP datagram? Explain how you determined the number of payload bytes.

    **R:** There are 20 bytes in the IP header, and 56 bytes as total length. The payload is the total length minus the header length, so 56 - 20 = 36 bytes.

4. Has this IP datagram been fragmented? Explain how you determined whether or not the datagram has been fragmented?

    **R:** No, because the fragment offset is 0.

5. Which fields in the IP datagram always change from one datagram to the next within this series of ICMP messages sent by your computer?

    **R:** The fields that change are the identification, the time to live and the header checksum

6. Which fields stay constant? Which of the fields must stay constant? Which fields must change? Why?

    **R:** The fields that stay constant are the version, protocol, header length, source ip (if we are sending from the same source), the destination ip (since we are sending to same destination), diffentiated Services. The fields that must stay constant are the previously mentioned, and the fields that must change are the identification, the time to live and the header checksum.

7. Describe the pattern you see in the values in the Identification field of the IP datagram?

    **R:** The pattern is that the identification increases by 1 in each packet.

8. What is the value in the Identification field and the TTL field?

    **R:** At first ICMP packet, the identification is 0x9302 and the TTL is 255.

9. Do these values remain unchanged for all of the ICMP TTL-exceeded replies sent to your computer by the nearest (first hop) router? Why?

    **R:** The identification changes for each packet because it is used to identify the fragments of a single datagram. The TTL changes because it is decremented by each hop.

## Fragmentation

10. Find the first ICMP Echo Request message that was sent by your computer after you changed the Packet Size in pingplotter to be 2000. Has that message been fragmented across more than one IP datagram?

    **R:** Yes, it has been fragmented.

11. Print out the first fragment of the fragmented IP datagram. What information in the IP header indicates that the datagram been fragmented? What information in the IP header indicates whether this is the first fragment versus a latter fragment? How long is this IP datagram?

    **R:** The flag of Fragment offset indicates that the datagram has been fragmented. The first datagram has a total length of 1500 bytes and the fragment offset is 0.

12. Print out the second fragment of the fragmented IP datagram. What information in the IP header indicates that this is not the first datagram fragment? Are the more fragments? How can you tell?

    **R:** We can tell that is not the first fragment because the fragment offset is 1480. There are more fragments because the more fragments flag is set to 1 on IPv4 header.

13. What fields change in the IP header between the first and second fragment?

    **R:** total length, header checksum, flags and fragment offset.

14. How many fragments were created from the original datagram?

    **R:** 2 fragments.

15. What fields change in the IP header among the fragments?

    **R:** The modified fields in the IP header across all the packets include the fragment offset and checksum. In addition, the total length and flags differ between the first two packets and the last packet. The initial two packets have a total length of 1500, with the more fragments flag set to 1. Conversely, the last packet has a total length of 540, with the more fragments flag set to 0.
