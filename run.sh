
tmux send-keys -t autoplius.0 C-c;
tmux send-keys -t autoplius.0 C-l;
tmux send-keys -t autoplius.0 "tmux clear-history" ENTER

# tmux send-keys -t autoplius.0 "go run ./apps/03-hello-struct/" ENTER
# tmux send-keys -t autoplius.0 "go run ./apps/04-hello-api/" ENTER
# tmux send-keys -t autoplius.0 "go test ./apps/05-hello-db/" ENTER
tmux send-keys -t autoplius.0 "go run ./apps/06-ads-db/" ENTER
# tmux send-keys -t autoplius.0 "go run ./apps/07-gui/" ENTER
# tmux send-keys -t autoplius.0 "go test -v ." ENTER

