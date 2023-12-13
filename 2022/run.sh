# If no parameters, show help
if [ $# -eq 0 ]; then
    echo "Runs AoC 2022 day's solution. First - tests. If they don't fail, then the solution itself."
    echo 
    echo "Usage: $0 <day>"
    echo
    echo "Example: $0 day01"
    exit 1
fi

# Get the day
day=$1

if [ ! -d $day ]; then
    echo "Directory $day does not exist"
    exit 1
fi

cd $day
go test . && go run .