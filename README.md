# ASCII-ART-WEB

```bash
             _____    _____   _____   _____                       _____    _______           __          __  ______   ____   
    /\      / ____|  / ____| |_   _| |_   _|              /\     |  __ \  |__   __|          \ \        / / |  ____| |  _ \  
   /  \    | (___   | |        | |     | |    ______     /  \    | |__) |    | |     ______   \ \  /\  / /  | |__    | |_) | 
  / /\ \    \___ \  | |        | |     | |   |______|   / /\ \   |  _  /     | |    |______|   \ \/  \/ /   |  __|   |  _ <  
 / ____ \   ____) | | |____   _| |_   _| |_            / ____ \  | | \ \     | |                \  /\  /    | |____  | |_) | 
/_/    \_\ |_____/   \_____| |_____| |_____|          /_/    \_\ |_|  \_\    |_|                 \/  \/     |______| |____/  
                                                                                                                             
```
Ascii-Art-Web allows users to interact with ASCII art generation through a web interface. The application provides options to select different banner styles (`shadow`, `standard`, and `thinkertoy`) and input text to be converted into ASCII art. The results are then displayed on the web page.

## Usage

### Running the Server

To run the server, follow these steps:

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/akothouma/ASCII_ART-WEB.git
      
    cd ascii-art-web
    ```

2. **Build and Run the Server:**

    Make sure you have [Go](https://go.dev/doc/install) installed on your machine. Then, run the following command to start the server:

    ```bash
    go run main.go
    ```

3. **Access the Web Application:**

    Open your web browser and navigate to `http://localhost:8080` to access the application.

### HTTP Endpoints

- **GET /**: Serves the HTML response for the main page.
- **POST /ascii-art**: Accepts text and banner style from the form and returns the ASCII art result.

### Error Handling

- **200 OK**: If the request is successful and processed without errors.
- **404 Not Found**: If the requested resource (template or banner) is not found.
- **400 Bad Request**: If the request is invalid or contains incorrect data.
- **500 Internal Server Error**: For unhandled errors or server issues.

## Authors

- [Akoth](https://github.com/akothouma)
- [Namayi](https://github.com/fnamayi) 