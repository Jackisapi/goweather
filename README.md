# GoWeather
## A CLI to fetch weather information written in GoLang

## Instalation

1. Install Golang
2. Make an account on https://openweathermap.org
3. Generate a API Key
4. Set the WEATHER_KEY ENV Variable
   ```bash
   export WEATHER_KEY=YOUR_KEY_HERE
   ```
5. Open the project directory and run 
    ```bash
    go build -o goweather
   ```
6. Finally try fetching the weather
    ```bash
   ./goweather -city=Boise -country=US
   ``` 


### Instalation (NixOS)
1. Make an account on https://openweathermap.org
2. Generate an API key
3. Open the project directory and edit the last Export in the shell.nix
    ```bash
    export WEATHER_KEY=YOUR_KEY_HERE
    ```
4. Save the file then run
    ```bash
   nix-shell
    ```
5. Finally try fetching the weather
    ```bash
    ./goweather -city=Boise -country=US
   ```
