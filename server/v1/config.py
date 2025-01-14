"""Module containing the project settings to configure various runtime environments.

This module defines settings for different runtime environments (development,
production, and staging) for the BurzCommerce API. It includes base settings as well as
environment-specific configurations.

The module relies on Pydantic's BaseSettings to define a structured way to configure
the API settings.

Key components of the module:
- Settings classes:
    - Settings: A base settings class that holds common configuration values like the
      API title, version, contact information, terms of service, etc.
    - DevelopmentSettings: A subclass of Settings containing settings specific to
      the development environment, such as enabling debug mode, specifying static file
      directories, and defining origins.
    - ProductionSettings: A subclass of Settings for production-specific
      configurations (currently empty, but may be extended as needed).
    - StagingSettings: A subclass of Settings for staging environment configurations
      (currently empty, but may be extended as needed).

- Constants:
    - _TITLE, _VERSION, _CONTACT, _DESCRIPTION, etc. define default values for
      common settings used across environments.

- get_settings() function:
    - A function that reads the ENVIRONMENT environment variable and returns the
      appropriate settings dictionary for the corresponding environment (development,
      staging, or production).
    - The function ensures that if no valid environment variable is found, it defaults
      to returning the DevelopmentSettings.
    - The returned dictionary contains the configuration settings, which can include
      strings, booleans, paths, or lists of key-value pairs.

The module provides a central place for managing settings that can vary depending on the
environment in which the API is running, helping ensure flexibility and consistency in
configuration management.
"""

import os
import pathlib
from typing import Any, Literal

from pydantic_settings import BaseSettings

_TITLE = "BurzPress API"
_VERSION = "v0.1.0"
_CONTACT = {"name": "Contact", "email": "somraj.saha@weburz.com"}
_DESCRIPTION = pathlib.Path(pathlib.Path.cwd() / "api" / "API.md").read_text()
_SUMMARY = """The BurzPress server-side API service to serve the client-side services.
"""
_TERMS_OF_SERVICE = "https://remington.bg/terms-of-service"
_TAGS_METADATA: list[dict[str, str]] = []


class Settings(BaseSettings):
    """Base settings class to be inherited across all runtime environments.

    This class defines common configuration values that are shared across all
    runtime environments (development, production, and staging). The settings
    include basic API information (e.g., title, version, contact details),
    API documentation URLs, and metadata for tags used in the API documentation.

    Attributes:
        title (str): The title of the API.
        version (str): The version of the API.
        contact (dict[str, str]): The contact details (name and email) for the API.
        description (str): A description of the API, loaded from an external file.
        summary (str): A brief summary of the API's purpose.
        terms_of_service (str): URL for the API's terms of service.
        tags_metadata (list[dict[str, str]]): Metadata for API tags used in the docs.
        openapi_url (str | None): URL for the OpenAPI schema.
        docs_url (str | None): URL for the API documentation.
        redoc_url (str | None): URL for the Redoc-generated API documentation.
    """

    title: str = _TITLE
    version: str = _VERSION
    contact: dict[str, str] = _CONTACT
    description: str = _DESCRIPTION
    summary: str = _SUMMARY
    terms_of_service: str = _TERMS_OF_SERVICE
    tags_metadata: list[dict[str, str]] = _TAGS_METADATA
    openapi_url: str | None = None
    docs_url: str | None = None
    redoc_url: str | None = None


class DevelopmentSettings(Settings):
    """Settings specific to the development runtime environment.

    This class extends the base Settings class and provides additional configurations
    relevant for a development environment, such as enabling debug mode, specifying
    paths for static files and templates, and defining the API documentation URLs. The
    development environment is typically used for local development and testing.

    Attributes:
        debug (bool): Flag to enable debug mode.
        static_files_dir (pathlib.Path): Path to the directory containing static files.
        templates_dir (pathlib.Path): Path to the directory containing templates.
        openapi_url (str | None): URL for the OpenAPI schema (like "/openapi.json".
        redoc_url (str | None): URL for the Redoc API documentation, Defaults to None.
        docs_url (str | None): URL for the API documentation, set to None by default.
        origins (list[str | Literal[""]]): List of allowed origins for CORS, typically
            allowing localhost for development.
    """

    debug: bool = True
    static_files_dir: pathlib.Path = pathlib.Path(pathlib.Path.cwd() / "api" / "static")
    templates_dir: pathlib.Path = pathlib.Path(pathlib.Path.cwd() / "api" / "templates")
    openapi_url: str | None = "/openapi.json"
    redoc_url: str | None = None
    docs_url: str | None = None
    origins: list[str | Literal[""]] = ["http://localhost:3000"]


class ProductionSettings(Settings):
    """Settings specific to the production runtime environment.

    This class extends the base Settings class and is intended to be used in the
    production environment. It can be customized with configurations that are specific
    to a live, production environment, such as database connections, logging, security
    settings, and any environment-specific features.

    This class is currently empty but can be extended to include production-specific
    settings as required.

    Attributes:
        Inherits all attributes from the base Settings class.
    """


class StagingSettings(Settings):
    """Settings specific to the staging runtime environment.

    This class extends the base Settings class and is used in the staging
    environment, which typically mirrors the production environment for
    pre-production testing. It can include configurations such as logging,
    database connections, or service endpoints that are specific to the staging
    environment.

    This class is currently empty but can be extended to include staging-specific
    settings as needed.

    Attributes:
        Inherits all attributes from the base Settings class.
    """


def get_settings() -> dict[str, Any]:
    """Fetches the appropriate settings object based on the runtime environment.

    The function checks the value of the ENVIRONMENT environment variable
    and returns the corresponding settings object as a dictionary (via .model_dump()).
    If no environment variable is set or if the value does not match a recognized
    environment, it defaults to returning the settings for the DevelopmentSettings
    environment.

    The returned dictionary contains configuration settings, which may include:
    - str: For string-type settings such as title, version, description, etc.
    - bool: For boolean-type settings like debug.
    - pathlib.Path: For path-type settings such as static_files_dir.
    - list[dict[str, str]]: For lists of key-value pairs, such as tags_metadata.

    Args:
        None

    Returns:
        dict[str, Union[str, bool, pathlib.Path, list[dict[str, str]]]]:
            A dictionary containing the settings for the selected environment. The
            dictionary's keys are setting names (e.g., title, version, etc.) and the
            values are the corresponding setting values, which can be strings, booleans,
            paths, or lists of dictionaries.

    Raises:
        None: This function does not raise any exceptions. It will return the settings
        as a dictionary even if the ENVIRONMENT value is not recognized, defaulting to
        DevelopmentSettings in that case.
    """
    match os.getenv("ENVIRONMENT"):
        case "staging":
            return StagingSettings().model_dump()
        case "production":
            return ProductionSettings().model_dump()
        case _:
            return DevelopmentSettings().model_dump()
