# Laboratório de Redes - Aula 6

## Capturing a bulk TCP transfer from your computer to a remote server - A first look at the captured trace

1. What is the IP address and TCP port number used by the client computer (source) that is transferring the file to gaia.cs.umass.edu? To answer this question, it’s probably easiest to select an HTTP message and explore the details of the TCP packet used to carry this HTTP message, using the “details of the selected packet header window” (refer to Figure 2 in the “Getting Started with Wireshark” Lab if you’re uncertain about the Wireshark windows).

   **R:** The IP is 192.168.1.102 and the PORT is 1161

2. What is the IP address of gaia.cs.umass.edu? On what port number is it sending and receiving TCP segments for this connection?

   **R:** The IP is 128.119.245.12 and the PORT which is sending is 80 and the PORT which is receiving is 1161

3. What is the IP address and TCP port number used by your client computer (source) to transfer the file to gaia.cs.umass.edu?

    **R:** The IP is XXX.XXX.XXX.XXX and the PORT is XXXX

## TCP Basics

4. What is the sequence number of the TCP SYN segment that is used to initiate the TCP connection between the client computer and gaia.cs.umass.edu? What is it in the segment that identifies the segment as a SYN segment?

    **R:** The sequence number is 0 and the SYN flag is 0x0002

5. What is the sequence number of the SYNACK segment sent by gaia.cs.umass.edu to the client computer in reply to the SYN? What is the value of the Acknowledgement field in the SYNACK segment? How did gaia.cs.umass.edu determine that value? What is it in the segment that identifies the segment as a SYNACK segment?

    **R:** The sequence number is 0 and the ACK is 1. The server determined the value of the ACK by adding 1 to the sequence number of the SYN. The SYNACK segment is identified by the flags SYN and ACK.

6. What is the sequence number of the TCP segment containing the HTTP POST command? Note that in order to find the POST command, you’ll need to dig into the packet content field at the bottom of the Wireshark window, looking for a segment with a “POST” within its DATA field.

    **R:** The sequence number is 164091

7. Consider the TCP segment containing the HTTP POST as the first segment in the TCP connection. What are the sequence numbers of the first six segments in the TCP connection (including the segment containing the HTTP POST)? At what time was each segment sent? When was the ACK for each segment received? Given the difference between when each TCP segment was sent, and when its acknowledgement was received, what is the RTT value for each of the six segments? What is the EstimatedRTT value (see Section 3.5.3, page 242 in text) after the receipt of each ACK? Assume that the value of the EstimatedRTT is equal to the measured RTT for the first segment, and then is computed using the EstimatedRTT equation on page 242 for all subsequent segments. **Note:** Wireshark has a nice feature that allows you to plot the RTT for each of the TCP segments sent. Select a TCP segment in the “listing of captured packets” window that is being sent from the client to the gaia.cs.umass edu server. Then select: Statistics->TCP Stream Graph->Round Trip Time Grap

    **R:** Question Jumped

8. What is the length of each of the first six TCP segments?

    **R:** The length of each TCP segments after HTTP POST (included) in order are: 565, 1460, 1460, 1460, 1460, 1460

9. What is the minimum amount of available buffer space advertised at the received for the entire trace? Does the lack of receiver buffer space ever throttle the sender?

    **R:** The minimum amount of availabe buffer space in the entire trace is 5840 bytes. The lack of receiver didn't throttle the sender.

10. Are there any retransmitted segments in the trace file? What did you check for (in the trace) in order to answer this question?

    **R:** Yes there are retransmitted segments. I realized that by checking to repeated sequence numbers.

10. How much data does the receiver typically acknowledge in an ACK? Can you identify cases where the receiver is ACKing every other received segment (see Table 3.2 on page 250 in the text).

    **R:** The ACK numbers increase in the sequence 566,2026,3486,4946 ... the ACK numbers increase by 1460 each time, indicating that the receiver is acknowledging 1460 bytes. There are cases where the receiver is ACKing every other received segment.

11. What is the throughput (bytes transferred per unit time) for the TCP connection? Explain how you calculated this value

    **R:** The last ACK number is 164091, as the first ack is 1, the total bytes transferred is 164090. The transmission time of the whole file is the difference of the time instant of the first TCP segment (0.026477 second for No.4 segment) and the time instant of the last ACK (5.455830 second for No. 202 segment), then, the total time is 5.455830 - 0.026477 = 5.429353 seconds. The throughput is 164090/5.429353 = 30200 bytes/second.
