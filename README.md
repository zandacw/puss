# Puss

Puss is a simple command-line utility that integrates GNU Pass and a fuzzy finder to streamline copying passwords and retrieving OTP codes. It allows you to quickly find and copy passwords stored in your `.password-store` directory and optionally retrieve OTP codes if available.

## Features

- List and search passwords using fuzzy finder (similar to fzf) 
- Copy selected password to clipboard
- Display OTP code if available

## Requirements

- [GNU Pass](https://www.passwordstore.org)
- [Pass OTP](https://github.com/tadfisher/pass-otp) (not required)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/zandacw/puss.git
   cd puss

2. Install

    ```bash
    go install

## Usage

1. Ensure your .password-store directory is properly set up and contains your encrypted passwords.

2. Run the command

   ```bash
   puss

3. Use the Fuzzy Finder to select a password. The selected password will be copied to your clipboard, and if an OTP is available, it will be displayed.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
