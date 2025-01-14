"""The entrypoint script for the "uvicorn" web server.

This script serves as the entry point for the FastAPI application, initializing the web
server (via Uvicorn) and setting up the necessary configurations, middleware, and routes
for the Remington API service. It loads the environment-specific settings using the
get_settings function and configures the FastAPI application accordingly.

Key features of this script:
- Settings Configuration: Fetches the environment-specific settings (e.g., debug
  mode, API title, CORS origins) using get_settings().
- FastAPI Initialization: Initializes the FastAPI application with settings like
  title, version, summary, contact details, and OpenAPI metadata.
- CORS Middleware: Configures Cross-Origin Resource Sharing (CORS) middleware to
  allow specified origins for API requests.
- Static Files: In development mode, it mounts static assets (e.g., for API
  documentation) and serves them via FastAPI.
- Custom API Documentation: Sets up a custom version of RapiDoc UI for displaying
  the OpenAPI documentation at /docs.
- Routing: Includes routers for various API endpoints such as health checks,
  product-related routes, and authentication.
- Uvicorn Configuration: Starts the Uvicorn web server with the appropriate host,
  port, logging, and configuration settings.

Environment Variables:
- ENVIRONMENT: Used to determine the runtime environment (e.g., "development",
  "production", "staging").
- Other settings like static_files_dir, templates_dir, host, port, and
  logging settings are loaded from the configuration.

This script is executed when starting the FastAPI application with Uvicorn, either for
local development (with live reloading) or in a production environment (with specific
configurations).
"""

import uvicorn
from fastapi import FastAPI

from config import get_settings

app = FastAPI()


@app.get("/")
def hello() -> dict[str, str]:
    """Lorem Ipsum."""
    return {"message": "Hello World!"}


settings = get_settings()


def main() -> None:
    """Invoke this function as the entrypoint of the script."""
    uvicorn.run(
        "api.main:app",
        port=settings.get("port", 8000),
        host=settings.get("host", "localhost"),
        reload=True,
    )


if __name__ == "__main__":
    main()
