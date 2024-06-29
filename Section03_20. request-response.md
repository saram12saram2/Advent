1. User and Their Computer:
   O
  /|\
  / \
 [____]   (Computer)


3. Typing in a URL:
 [www.example.com]
   O
  /|\
  / \
 [____]   (Computer)


5. DNS Lookup:
 [www.example.com]
   O
  /|\
  / \
 [____]   (Computer)
    |
    V
  [DNS]
    |
   IP: 192.168.1.1
   

7. Sending a Request:
   O
  /|\
  / \
 [____] -----> [Server]


8. Server-Side Processing:
   O
  /|\
  / \
 [____] -----> [Server]
                /  |  \
             [DB] [Cache] [App]


9. Static Content Handling:
   O
  /|\
  / \
 [____] <----- [Server]
      HTML file


10. Dynamic Content Handling:
   O
  /|\
  / \
 [____] <----- [Server]
                /  |  \
             [DB]<-> [Cache]<-> [App]


11. Response:
   O
  /|\
  / \
 [____] <----- [Server]
              HTML, CSS, JS


12. Cookies and Authentication:
   O
  /|\
  / \
 [____] <-----> [Server]
  Cookie <----> Cookie


13. Browser Rendering:
   O
  /|\
  / \
 [____]
 [Web Page]


14. Monolithic vs. Microservices:
[Monolithic Server]

[Auth] - [Billing] - [Content]
(Microservices)
