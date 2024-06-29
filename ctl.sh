#!/bin/bash

APP_NAME="ecom"
MAIN_PACKAGE="app/main.go"
MAIN_FILE="main.go"
BUILD_OUTPUT="bin/$APP_NAME"

build_app() {
    echo "Building $APP_NAME..."
    go build -o $BUILD_OUTPUT $MAIN_PACKAGE
    if [ $? -eq 0 ]; then
        echo "Build successful."
    else
        echo "Build failed."
    fi
}

run_app() {
    echo "Running $APP_NAME..."
    ./$BUILD_OUTPUT
}

clean() {
    echo "Cleaning up..."
    go clean
    rm -rf $BUILD_OUTPUT
    echo "Clean up complete."
}

main() {
    case "$1" in
        build)
            build_app
            ;;
        run)
            run_app
            ;;
        clean)
            clean
            ;;
        *)
            echo "Usage: $0 {build|run|clean}"
            exit 1
            ;;
    esac
}

main "$@"
