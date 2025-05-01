# Test provider based on Terraform Plugin SDKv2

This repository demonstrates the implementation of a custom Terraform provider using the Terraform Plugin SDKv2.  
The provider interacts with a simple SQLite database to manage "person" resources, showcasing key features of Terraform
provider development.

## Features

- **CRUD Operations**: Full support for Create, Read, Update, and Delete operations on `person` resources.
- **Drift Detection**: Automatically detects and reconciles changes made outside of Terraform to ensure the state
  remains consistent.
- **Resource Recreation**: Supports resource recreation when required, such as when a resource is forcefully replaced.
- **Import Functionality**: Allows importing existing resources into Terraform state for management.

This project is designed for learning purposes and provides a hands-on example of how to build and test a Terraform
provider.
