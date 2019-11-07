if [ "$1" != "-b" ] || [ -z $"$2" ]; then
  echo "No branch defined. Use flag -b to define the branch to pull from"
else
  git pull origin $"2"
  easyjson --all ./internal/pkg/models
  ./scripts/run.sh
fi
