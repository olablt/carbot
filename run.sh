
tmux send-keys -t autoplius.0 C-c;
tmux send-keys -t autoplius.0 C-l;
tmux send-keys -t autoplius.0 "tmux clear-history" ENTER

tmux send-keys -t autoplius.0 "go run ./apps/03-hello-struct/" ENTER

