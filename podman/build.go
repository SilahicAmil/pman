package podman

// Run the actual Podman API build commands

// Check if image exists?
// IF does prompt use if they wanna override/continue?

// build tmp container
// - POST /v6.0.0/libpod/containers/create

// copy any context
// - PUT /v6.0.0/libpod/containers/{container}/archive?path=/app

// commit the container
// - POST /v6.0.0/libpod/commit?container={container}&repo={image}&tag=latest

// delete tmp container
// - DELETE /v6.0.0/libpod/containers/{container}

// --- PMAN UP BELOW FOR THE FUTURE --- //

// create container
// - POST /v6.0.0/libpod/containers/create

// start container
// -POST /v6.0.0/libpod/containers/{container}/start
