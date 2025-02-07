prepare_directory_to_test () {
    cd ..
    go build
    cd e2e
    mkdir buckmate-test-directory
    mv ../buckmate buckmate-test-directory/buckmate-executable
    cd buckmate-test-directory
}

compare_results () {
    if diff -r "$1" "$2" --exclude=".gitignore"; then
        echo "SUCCESS"
        cd ../
        rm -rf buckmate-test-directory
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
    compare_results "../result" "buckmate-target"
}

s3_to_s3 () {
    echo "Starting s3-to-s3 test"
    prepare_directory_to_test
    cp -R ../../example/s3-to-s3/* ./
    ./buckmate-executable --path buckmate apply
    mkdir buckmate-target
    aws s3 cp s3://buckmate-target buckmate-target --recursive
    compare_results "../result" "buckmate-target"
}

local_to_s3 () {
    echo "Starting local-to-s3 test"
    prepare_directory_to_test
    cp -R ../../example/local-to-s3/* ./
    ./buckmate-executable --path buckmate apply
    mkdir buckmate-target
    aws s3 cp s3://buckmate-target buckmate-target --recursive
    compare_results "../result" "buckmate-target"
}

s3_to_local () {
    echo "Starting s3-to-local test"
    prepare_directory_to_test
    cp -R ../../example/s3-to-local/* ./
    ./buckmate-executable --path buckmate apply
    compare_results "../result" "buckmate-target"
}

dry_local () {
    echo "Starting dry local run test"
    prepare_directory_to_test
    cp -R ../../example/local-to-local/* ./
    result=`./buckmate-executable --path buckmate apply --dry 2>&1`
    var_path=$(echo "$result" | awk '{for(i=1;i<=NF;i++) if($i ~ /^\/var\//) print $i}')
    compare_results "../result" "$var_path"
}

dry_remote () {
    echo "Starting dry remote run test"
    prepare_directory_to_test
    cp -R ../../example/local-to-s3/* ./
    result=`./buckmate-executable --path buckmate apply --dry 2>&1`
    var_path=$(echo "$result" | awk '{for(i=1;i<=NF;i++) if($i ~ /^\/var\//) print $i}')
    compare_results "../result" "$var_path"
}
#
cache_control_metadata () {
    echo "Starting local-to-s3-cache-metadata test"
    prepare_directory_to_test
    cp -R ../../example/local-to-s3-cache-metadata/* ./
    ./buckmate-executable --path buckmate apply
    aws s3api head-object --bucket buckmate-target --key index.html >> tmp1
    aws s3api head-object --bucket buckmate-target --key common-file.json >> tmp2
    tmp1CacheControl=$(cat tmp1 | jq '.CacheControl')
    tmp1MetadataKey=$(cat tmp1 | jq '.Metadata."some-metadata-key"')
    tmp2CacheControl=$(cat tmp2 | jq '.CacheControl')
    tmp2MetadataKey=$(cat tmp2 | jq '.Metadata."some-metadata-key"')
    if  [ "$tmp1CacheControl" != "\"no-cache\"" ]; then
        echo "FAILURE1"
        exit 1
    fi
    if  [ "$tmp1MetadataKey" != "\"some-metadata-value\"" ]; then
        echo "FAILURE2"
        exit 1
    fi
    if  [ "$tmp2CacheControl" != "null" ]; then
        echo "FAILURE3"
        exit 1
    fi
    if  [ "$tmp2MetadataKey" != "null" ]; then
        echo "FAILURE4"
        exit 1
    fi
    cd ../
    rm -r buckmate-test-directory
    echo "SUCCESS"
}

keep_previous() {
    echo "Starting local-to-s3-keep-previous test"
    prepare_directory_to_test
    cp -R ../../example/local-to-s3-keep-previous/* ./
    ./buckmate-executable --path buckmate apply
    mkdir buckmate-target
    aws s3 cp s3://buckmate-keep-previous buckmate-target --recursive
    compare_results "../result-keep-previous" "buckmate-target"
}

local_to_local
s3_to_s3
local_to_s3
s3_to_local
dry_local
dry_remote
cache_control_metadata
keep_previous