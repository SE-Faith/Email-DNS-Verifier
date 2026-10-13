# Email DNS Verifier

A small Go command-line tool that checks whether an email address is syntactically valid and inspects the domain's DNS records for common mail-related signals.

This project is designed for quick checks of:
- email format validation
- domain IP lookup
- MX record lookup
- SPF record lookup
- DMARC record lookup
- table-driven test coverage

> This tool does not verify that a mailbox exists, that an email will be delivered, or that a domain is fully configured for email sending.

## Project overview

The program reads an email address from standard input, validates the address format, extracts the domain portion, and then checks the domain using Go's standard library networking and DNS functions.

It is useful as a lightweight diagnostic tool when you want to confirm that a domain appears to have the basic DNS records commonly associated with email.

## Features

- Email format validation using a practical sanity check for local-part and domain structure
- IP lookup for the domain using `net.LookupIP`
- MX lookup for the domain using `net.LookupMX`
- SPF record detection using `net.LookupTXT` and checking for `v=spf1`
- DMARC record detection using `net.LookupTXT` on `_dmarc.<domain>` and checking for `v=dmarc1`
- Table-driven unit tests covering valid and invalid email inputs

## Project structure

```text
Email-DNS-verifier/
├── go.mod
├── main.go
├── main_test.go
└── README.md
```

## Technologies used

- Go
- Standard library packages:
  - `bufio`
  - `fmt`
  - `net`
  - `os`
  - `strings`
- DNS lookups via `net.LookupIP`, `net.LookupMX`, and `net.LookupTXT`

## Prerequisites

Before running the project, make sure you have:

- Go installed and available on your `PATH`
- A working internet connection for live DNS lookups

## Installation

From the project folder:

```bash
cd Email-DNS-verifier
```



## Usage

Run the tool:

```bash
go run .
```

When prompted, enter an email address, for example:

```text
Enter an email address:
user@gmail.com
```

You can also pipe input directly:

```bash
printf "user@gmail.com\n" | go run .
```

## Example output

```text
Enter an email address:
user@gmail.com
Valid email address
Domain IP Lookup: Found 1 IP address(es)
Domain MX Lookup: Found 5 MX records
Domain SPF Lookup: Found SPF record: v=spf1 include:_spf.google.com ~all
Domain DMARC Lookup: Found DMARC record: v=DMARC1; p=none; sp=quarantine; rua=mailto:mailauth-reports@google.com;
```

The exact output depends on the domain and the live DNS records returned by your resolver.

## Testing

The project includes table-driven tests for email validation in `main_test.go`.

Run the tests with:

```bash
go test -v
```

This runs the validation cases and confirms the expected behavior for valid and invalid email addresses.

## Limitations

- The program only performs a basic email format check; it is not a full email validation engine.
- It checks for the presence of DNS records, but it does not confirm whether a mailbox actually exists.
- It does not guarantee email deliverability or inbox acceptance.
- DNS results depend on the current network and resolver environment.
- SPF and DMARC checks only look for the presence of records, not whether they are correct for every sending configuration.




