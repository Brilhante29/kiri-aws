# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.25.x  | :white_check_mark: |
| < 0.25  | :x:                |

## Reporting a Vulnerability

We take security seriously. If you discover a security vulnerability in `kiri-aws`, please do **not** open a public GitHub issue.

Instead, please report vulnerabilities via email or GitHub Private Vulnerability Reporting:

- **Security Email:** `guilhermebrilhante00@gmail.com`
- **Response Time:** We aim to acknowledge receipt within 48 hours and provide an initial assessment within 5 business days.

## Security Hardening & Best Practices

- `kiri-aws` is designed as an **offline local AWS emulator** for development and testing.
- It does **not** validate live AWS IAM signatures or perform cloud authentication against AWS APIs.
- Do **not** expose `kiri-aws` directly to untrusted public networks without an authentication proxy or VPN.
