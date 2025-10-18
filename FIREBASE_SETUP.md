# Firebase Setup Guide

This guide explains how to configure the application to use Google Firebase (Firestore) as the database backend.

## Prerequisites

1. A Google Cloud account
2. A Google Cloud project with Firestore enabled
3. Service account credentials for authentication

## Step 1: Create a Google Cloud Project

1. Go to the [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Note your Project ID (you'll need this later)

## Step 2: Enable Firestore

1. In the Google Cloud Console, navigate to **Firestore**
2. Click **Create Database**
3. Choose **Native mode** (recommended)
4. Select a location for your database
5. Click **Create Database**

## Step 3: Create a Service Account

1. Navigate to **IAM & Admin** > **Service Accounts**
2. Click **Create Service Account**
3. Provide a name and description
4. Grant the role **Cloud Datastore User** (or **Firestore Service Agent**)
5. Click **Done**

## Step 4: Generate Service Account Key

1. Click on the service account you just created
2. Go to the **Keys** tab
3. Click **Add Key** > **Create new key**
4. Choose **JSON** format
5. Click **Create** - a JSON key file will be downloaded

**Important**: Keep this file secure and never commit it to version control!

## Step 5: Configure Application

There are two ways to configure the application to use Firebase:

### Option 1: Environment Variables (Recommended)

Set the following environment variables:

```bash
# Set the database type to Firebase
export DB_TYPE=firebase

# Set your Google Cloud project ID
export FIREBASE_PROJECT_ID=your-project-id

# Optional: Set custom collection name (default is "books")
export FIREBASE_COLLECTION=books

# Set the path to your service account key file
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/your-service-account-key.json
```

### Option 2: Environment File

Create a `.env` file in the project root (add it to `.gitignore`):

```bash
DB_TYPE=firebase
FIREBASE_PROJECT_ID=your-project-id
FIREBASE_COLLECTION=books
GOOGLE_APPLICATION_CREDENTIALS=/path/to/your-service-account-key.json
```

## Step 6: Run the Application

```bash
# Build the application
go build -o fiber-app

# Run with Firebase
./fiber-app
```

You should see:
```
Initializing Firebase repository with project: your-project-id, collection: books
```

## Switching Between Databases

The application supports easy switching between database backends:

### Use In-Memory Database (Default)
```bash
export DB_TYPE=memory
./fiber-app
```

### Use Firebase Database
```bash
export DB_TYPE=firebase
export FIREBASE_PROJECT_ID=your-project-id
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json
./fiber-app
```

## Data Structure in Firestore

Books are stored in Firestore with the following structure:

```
Collection: books (or your custom collection name)
Document ID: Auto-generated unique ID
Fields:
  - id: string (same as document ID)
  - title: string
  - author: string
  - isbn: string
  - year: number
```

## Security Considerations

1. **Never commit credentials**: Add `*.json` service account files to `.gitignore`
2. **Use IAM roles**: Grant minimum necessary permissions
3. **Firestore Security Rules**: Configure appropriate security rules in the Firebase Console
4. **Environment-specific configs**: Use different service accounts for development, staging, and production

## Example Firestore Security Rules

```javascript
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /books/{bookId} {
      // Allow read access to all users
      allow read: if true;
      
      // Allow write access only to authenticated users
      // Adjust based on your authentication setup
      allow write: if request.auth != null;
    }
  }
}
```

## Troubleshooting

### Error: "Failed to initialize Firebase repository"

- Verify your Project ID is correct
- Ensure Firestore is enabled in your project
- Check that `GOOGLE_APPLICATION_CREDENTIALS` points to a valid JSON key file

### Error: "Permission denied"

- Verify the service account has the correct IAM roles
- Check Firestore security rules
- Ensure the service account key is not expired

### Error: "Collection not found"

- Firestore collections are created automatically when you add the first document
- Try creating a book via the API - the collection will be created automatically

## Testing Firebase Integration

1. Start the application with Firebase configuration
2. Create a book:
   ```bash
   curl -X POST http://localhost:3000/api/books \
     -H "Content-Type: application/json" \
     -d '{"title":"Test Book","author":"Test Author","isbn":"123","year":2024}'
   ```
3. Check the Firestore console to verify the book was created
4. Retrieve books via the API:
   ```bash
   curl http://localhost:3000/api/books
   ```

## Additional Resources

- [Firestore Documentation](https://cloud.google.com/firestore/docs)
- [Go Client Library](https://pkg.go.dev/cloud.google.com/go/firestore)
- [Authentication Guide](https://cloud.google.com/docs/authentication/getting-started)
