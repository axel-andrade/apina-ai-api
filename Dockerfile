# Use an official Python runtime as a parent image
FROM python:3.8-slim

# Set the working directory
WORKDIR /app

# Install necessary packages
RUN apt-get update && apt-get install -y \
    wget \
    curl \
    git \
    build-essential \
    ffmpeg \
    && rm -rf /var/lib/apt/lists/*

# Install DeepSpeech
RUN pip install deepspeech

# Download pre-trained English model files
RUN wget -q https://github.com/mozilla/DeepSpeech/releases/download/v0.9.3/deepspeech-0.9.3-models.pbmm \
    && wget -q https://github.com/mozilla/DeepSpeech/releases/download/v0.9.3/deepspeech-0.9.3-models.scorer

# Move the models to a known directory
RUN mkdir -p /opt/deepspeech \
    && mv deepspeech-0.9.3-models.pbmm /opt/deepspeech/ \
    && mv deepspeech-0.9.3-models.scorer /opt/deepspeech/

# Install Golang (if needed, adjust version)
RUN wget -q https://golang.org/dl/go1.22.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz \
    && rm go1.22.4.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="$GOPATH/bin:/usr/local/go/bin:$PATH"

# Install build essentials and other tools (equivalent to build-base in Alpine)
RUN apt-get update && apt-get install -y \
    build-essential

# Set environment variables for Go build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy only the go.mod and go.sum to allow caching of dependencies
COPY go.mod go.sum ./

# Download and install dependencies, including updating go.sum and fetching missing dependencies
RUN go mod tidy

# Copy the rest of the source code to the working directory
COPY . .

# Compile the Go application
RUN go build -o app

# Remove any unnecessary dependencies that might have been installed during the build
RUN apt-get remove -y build-essential && apt-get autoremove -y

# Add certificates for HTTPS support
RUN apt-get update && apt-get install -y ca-certificates

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./app"]
