package main

import (
	"flag"
	"log"

	"github.com/pimmerks/vault-s3-snapshot/auth"
	"github.com/pimmerks/vault-s3-snapshot/config"
	"github.com/pimmerks/vault-s3-snapshot/config/enums"
	"github.com/pimmerks/vault-s3-snapshot/writer"
)

func main() {
	configPath := flag.String("config", "", "The path to the config file (in json)")
	flag.Parse()

	if *configPath == "" {
		log.Fatalln("Please provide a config path (--config /path/to/config.json).")
	}

	log.Printf("Reading configuration from '%v'\n", *configPath)
	c, err := config.ReadConfig(*configPath)

	if err != nil {
		log.Fatalln("Configuration could not be found")
	}

	var authenticator auth.VaultAuthenticator
	switch c.VaultAuthMethod {
	case enums.Token:
		authenticator = auth.CreateTokenAuth(c)
	default:
		log.Fatalln("Please specify the Vault auth method")
	}

	log.Printf("Using authenticator '%v'\n", authenticator.GetType())

	// snapshotter, err := snapshot_agent.NewSnapshotter(c)
	// if err != nil {
	// 	log.Fatalln("Cannot instantiate snapshotter.", err)
	// }

	// if snapshotter.TokenExpiration.Before(time.Now()) {
	// 	switch c.VaultAuthMethod {
	// 	case "k8s":
	// 		snapshotter.SetClientTokenFromK8sAuth(c)
	// 	default:
	// 		snapshotter.SetClientTokenFromAppRole(c)
	// 	}
	// }
	// leader, err := snapshotter.API.Sys().Leader()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	log.Fatalln("Unable to determine leader instance.  The snapshot agent will only run on the leader node.  Are you running this daemon on a Vault instance?")
	// }
	// leaderIsSelf := leader.IsSelf
	// if !leaderIsSelf {
	// 	log.Println("Not running on leader node, skipping.")
	// }

	// var snapshot bytes.Buffer
	// err = snapshotter.API.Sys().RaftSnapshot(&snapshot)
	// if err != nil {
	// 	log.Fatalln("Unable to generate snapshot", err.Error())
	// }

	writers := []writer.SnapshotWriter{}

	if c.Local.Path != "" {
		writers = append(writers, writer.CreateLocalPathSnapshotWriter(c))
	}

	if c.AWS.Bucket != "" {
		writers = append(writers, writer.CreateS3SnapshotWriter(c))
	}

	// now := time.Now().UnixNano()
	for _, v := range writers {
		log.Printf("Starting snapshot writer %v\n", v.GetType())
		// v.WriteSnapshot(&snapshot, now)
	}
}

func logSnapshotError(dest, snapshotPath string, err error) {
	if err != nil {
		log.Printf("Failed to generate %s snapshot to %s: %v\n", dest, snapshotPath, err)
	} else {
		log.Printf("Successfully created %s snapshot to %s\n", dest, snapshotPath)
	}
}
