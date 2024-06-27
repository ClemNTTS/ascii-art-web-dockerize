Ascii-Art-Web

Ascii-Art-Web consists of creating and running a server that provides a web GUI (graphical user interface) version of our previous project, ascii-art. Users can input text and select different banners to generate ASCII art.

Features
Web GUI: A user-friendly web interface to generate ASCII art.
Banner Options: Support for multiple banners:
    shadow
    standard
    thinkertoy

HTTP Endpoints:
GET /: Main page that sends an HTML response.
POST /ascii-art: Sends data (text and banner) to the Go server.

How to use :
Write in the text area, a text that you want to transform in ascii-art and click on Submit button.