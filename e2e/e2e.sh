prepare_directory_to_test () {
    cd ..
    go build
    cd e2e
    mkdir buckmate-test-directory
    mv ../buckmate buckmate-test-directory/buckmate-executable
    cd buckmate-test-directory
}

compare_results () {
    if diff -r "$1" "$2"; then
        echo "SUCCESS"
        cd ../
        rm -rf buckmate-test-directory
        exit 0
    else
        echo "FAILURE"
        exit 1
    fi
}

local_to_local () {
    echo "Starting local-to-local test"
    prepare_directory_to_test
    cp -R ../../example/local-to-local/* ./
    ./buckmate-executable --path buckmate apply
    compare_results "../local-to-local-result" "buckmate-target"
}

local_to_local