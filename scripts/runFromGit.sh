if [ "$1" != "-b" ] || [ -z $"$2" ]; then
  echo "pulling from origin..."
  git pull origin
else
  echo "pulling from origin $2..."
  git pull origin "$2"
fi
./scripts/run.sh
