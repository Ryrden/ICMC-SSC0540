# Laboratório de Redes - Aula 4

## The Basic HTTP GET/response interaction - Questions

1. Is your browser running HTTP version 1.0 or 1.1? What version of HTTP is the
server running?

    **R:** Server is running HTTP 1.1, browser is running HTTP 1.0

2. What languages (if any) does your browser indicate that it can accept to the
server?

    **R:** Accept-Language: pt-Br, pt

3. What is the IP address of your computer? Of the gaia.cs.umass.edu server?

    **R:** XXX.XX.XXX.XXX //TODO: Put IP before submitting

4. What is the status code returned from the server to your browser?

    **R:** The status code returned from the server to my browser is 200

5. When was the HTML file that you are retrieving last modified at the server?

    **R:** Last-Modified: Fri, 14 Apr 2023 05:59:01 GMT

6. How many bytes of content are being returned to your browser?

    **R:** There are 128 bytes of content being returned to my browser

7. By inspecting the raw data in the packet content window, do you see any headers within the data that are not displayed in the packet-listing window? If so, name one.

    **R:** Yes, the header "Content-Type: text/html"

## The HTTP CONDITIONAL GET/response interaction - Questions

8. Inspect the contents of the first HTTP GET request from your browser to the server. Do you see an “IF-MODIFIED-SINCE” line in the HTTP GET?

    **R:** On the first HTTP GET request, there is no "IF-MODIFIED-SINCE" line, but on the second one, there is.

9. Inspect the contents of the server response. Did the server explicitly return the contents of the file? How can you tell?

    **R:** Yes, cause the 200 status code, another request is made to the server to get the file and the server explicitly returns the contents of the file.

10. Now inspect the contents of the second HTTP GET request from your browser to the server. Do you see an “IF-MODIFIED-SINCE:” line in the HTTP GET? If so, what information follows the “IF-MODIFIED-SINCE:” header?

    **R:** Yes, the information that follows the "IF-MODIFIED-SINCE" header is "Fri, 14 Apr 2023 05:59:01 GMT"

11. What is the HTTP status code and phrase returned from the server in response to this second HTTP GET? Did the server explicitly return the contents of the file? Explain

    **R:** No, cause the 304 status code, the server does not explicitly return the contents of the file.

## Retrieving Long Documents - Questions

12. How many HTTP GET request messages did your browser send? Which packet number in the trace contains the GET message for the Bill or Rights?

    **R:** There were 2 HTTP GET request messages, the packet number is 4 Reassembled TCP Segments

13. Which packet number in the trace contains the status code and phrase associated with the response to the HTTP GET request?

    **R:** 

14. What is the status code and phrase in the response?

    **R:**

15. How many data-containing TCP segments were needed to carry the single HTTP
response and the text of the Bill of Rights?

    **R:**

## HTML Documents with Embedded Objects - Questions

16. How many HTTP GET request messages did your browser send? To which Internet addresses were these GET requests sent?

    **R:** There were 4 HTTP GET request messages, the Internet addresses were

17. Can you tell whether your browser downloaded the two images serially, or whether they were downloaded from the two web sites in parallel? Explain

    **R:**

## HTTP Authentication - Questions

18. What is the server’s response (status code and phrase) in response to the initial
HTTP GET message from your browser?

    **R:** The server's response is 401 Unauthorized

19. When your browser’s sends the HTTP GET message for the second time, what
new field is included in the HTTP GET message?

    **R:** Last-Modified, ETag and Accept-Ranges headers
