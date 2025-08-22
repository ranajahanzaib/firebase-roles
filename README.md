# Firebase Roles CLI

[![GitHub license](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE) [![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-brightgreen.svg)](./CONTRIBUTING.md)

A standalone CLI tool to assign Firebase custom roles to users using a service account key.

## Features

- Assign roles like `admin` or custom roles to Firebase Authentication users.
- Works via service account credentials.
- Lightweight and secure; no server or cloud function required.
- Supports multiple roles per user (optional).

## Installation

Build the CLI from source (Go required):

```bash
git clone https://github.com/ranajahanzaib/firebase-roles.git
cd firebase-roles
go build -o firebase-roles main.go
```

### Usage

```
./firebase-roles --service-account=serviceAccountKey.json --uid=USER_UID --role=admin

  --service-account: Path to your Firebase service account JSON key.
  --uid: UID of the Firebase user to assign the role.
  --role: Role to assign (e.g., admin). Can be used multiple times if supported.

Example:

./firebase-roles --service-account=serviceAccountKey.json --uid=abc123 --role=admin --role=moderator
```

## Contributing

We’d love to accept your patches and contributions to this project. There are just a few guidelines you need to follow.

### [Code of Conduct](./CODE_OF_CONDUCT.md)

This project follows the [Contributor Covenant](https://www.contributor-covenant.org/) as its Code of Conduct, and we expect all project participants to adhere to it. Kindly read the [full guide](./CODE_OF_CONDUCT.md) to understand what actions will not be tolerated.

### [Contributing Guide](./CONTRIBUTING.md)

Read our [contributing guide](./CONTRIBUTING.md) to learn about our development process, how to propose bug fixes and improvements, and how to build and test your changes to the project.

### Issues

Feel free to submit issues and enhancement requests. Use the template provided when creating an issue to ensure your request is clear and actionable.

See the open issues for a complete list of proposed features (and known issues).

## [License](./LICENSE)

This project is licensed under the [MIT License](./LICENSE), meaning that you’re free to modify, distribute, and/or use it for any commercial or private project.
