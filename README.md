#TODO

- Feedback for filling in forms
- User verification

- Auth:
  - Secret for JWT
  - separate admin and user privileges
  - add user level and user specific restrictions to pages 
  - add token refresh and expiry check
  - update login pages (EB: make one unified login?)
  - consider JWT vs session tokens w/ db

  - Form Submissions
    - Note: Use JWT or session to store note content until note is submitted to prevent partial storage of duplicate info
    - Note: for each new page, if data has been posted use it to update the session note (if available)
