# sess

## What is sess?

**sess** is a simple library that implements some alternative to jwt. As simple and clear as possible. Only HMAC signature, which is based on a slice of bytes key.

## How to install

To install sess, run the command:

```
go get -u github.com/nyan2d/sess
```

Add the library to your code:

```
import "github.com/nyan2d/sess"
```

## Usage

### Creating a sess instance and sessions

To create an instance of **sess**, we need a secret key in the form of a slice of bytes. The best way to get a strong key is to use a random number generator. But in our example we will use a simple string.

```
secretKey := []byte("secret")
s := sess.New(secretKey)
```

Now that we have an instance of **sess**, we can create sessions for users.

```
userID := 5
validDuration := 5 * time.Second
session := s.CreateSession(userID, validDuration)
```

Next, the created session must be signed, after which no unwanted changes can be applied to it

```
s.SignSession(session)
```

### Session validation

To check the validity of a session, you just need to call the *IsSessionValid* function

```
s.IsSessionValid(session)
```

### Session Updating

An expired session can be extended using the *UpdateSession* function

```
validDuration := 5 * time.Secound
s.UpdateSession(session, validDuration)
```

When you do so, the sessions will be signed automatically.

### Interaction with an http server

Interaction with http requests is done by using cookies. We can read and write sessions to cookies.

```
func exampleHandler(w http.ResponseWriter, r *http.Request) {
    // Read session from cookies
    session, err := sess.ReadFromCookies(r)

    // Write the session to a cookies
    sess.WriteToCookies(w, session)
}
```

## Project Status

At the moment **sess** is used in several pet projects as well as in a small SPA at my work.

## License

This library is distributed under the **MIT** license.