#TODO

- Feedback for filling in forms
- User verification

- Auth:
  - Secret for JWT
  - separate admin and user privileges
  - add user level and user specific restrictions to pages (EB: this could be done through middleware, wrapping a route with the appropriate check for privileges)
  - add token refresh and expiry check
  - update login pages (EB: make one unified login?)