/*
 * Copyright 2021 Teo Mrnjavac <teo.mrnjavac@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package replication

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

/*
POOLNAME="${ALIZSYNC_POOL_NAME:-aliZsync}"
IMG_FILE_PATH="${ALIZSYNC_IMG_FILE_PATH:-$HOME/$POOLNAME.img}"
TRANSPORT="${ALIZSYNC_TRANSPORT:-mbuffer}"

INVENTORY_FILE="${ALIZSYNC_INVENTORY:-/etc/o2.d/aliZsync_inventory}"
TARGET_ROOT="${ALIZSYNC_TARGET_ROOT:-/opt/alizsw}"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
TAG_NAME="${2:-$TIMESTAMP}"
N_WORKERS="${ALIZSYNC_WORKERS:-10}"

 */
const (
	N_WORKERS = 10          // maximum total cluster-wide concurrent jobs
	BRANCHING_FACTOR = 3    // transfer jobs per source
)

type Transport string
const (
	Transport_NULL = Transport("")
	Transport_SSH = Transport("ssh")
	Transport_NETCAT = Transport("netcat")
	Transport_MBUFFER = Transport("mbuffer")
)

func Sync(source string, targets []string) {
	timestamp := time.Now()
	sources := make(chan string, len(targets) + 1)	// buffered channel through which we manage intermediate sources

	concurrentGoroutines := make(chan struct{}, N_WORKERS)

	sources <- source // we init the sources channel with the first source

	var wg sync.WaitGroup
	for i := 0; i < len(targets); i += BRANCHING_FACTOR {

		lastIndex := int(math.Min(float64(len(targets)), float64(i + BRANCHING_FACTOR)))	// make sure we don't overshoot

		source := <-sources
		currentTargets := targets[i:lastIndex]
		for _, target := range currentTargets {
			concurrentGoroutines <- struct{}{}
			wg.Add(1)
			target := target
			go func() {
				defer wg.Done()
				fmt.Println("replicating from", source)
				err := doSync(source, target)
				if err != nil {
					fmt.Println("error", target)
				} else {
					sources <- target	// means the target was successfully synced, so it can become a source
				}
				fmt.Println("\tfinished", target)
				<-concurrentGoroutines
			}()
		}
	}
	wg.Wait()

	elapsed := time.Since(timestamp)
	fmt.Printf("done in %s (%f seconds per job)", elapsed.String(), elapsed.Seconds() / float64(len(targets)))
}


func doSync(source string, target string) error {
	r := float64(rand.Intn(9) + 1) // dither from 1 to 10
	r = 3 + r/10 // from 3.1 to 4s random duration
	fmt.Printf("\tsyncing %s to %s for %f seconds\n", source, target, r)
	time.Sleep(time.Duration(r) * time.Second)
	return nil
}