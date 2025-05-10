system architecture 
![system_architecture](https://github.com/user-attachments/assets/84a800c3-ffdf-4923-814f-3c67e0aaa633)

## 🚀 Deploying to Heroku with Docker

### 🔧 Prerequisites

* Docker installed
* Heroku CLI installed and logged in

---

### 📆 Steps to Deploy

1. **Create a `Dockerfile`** in your project root.

2. **Build the Docker image locally (optional for testing)**:

   ```bash
   docker build -t your-image-name .
   ```

3. **Run the image locally for testing**:

   ```bash
   docker run -p 5000:5000 -e PORT=5000 your-image-name
   ```

4. **Login to Heroku Container Registry**:

   ```bash
   heroku container:login
   ```

5. **Create a new Heroku app**:

   ```bash
   heroku create
   ```

6. **Push the Docker container to Heroku**:

   ```bash
   heroku container:push web
   ```

7. **Release the container on Heroku**:

   ```bash
   heroku container:release web
   ```
