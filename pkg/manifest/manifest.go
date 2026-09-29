

package manifest

import (
	"fmt"
	"log"
	"strings"
	"crypto/sha256"
    "io"
    "os"
	"encoding/hex"
	"path/filepath"


	"encoding/json"
)



type ReleaseManifest struct {
Tag *string `json:"image_tag"`
ImageUrl *string `json:"image_url"`
ImageSha *string `json:"image_sha"`
ImageSize *uint64 `json:"image_size"`
}

func ParseManifest(data[]byte)(*ReleaseManifest, error)
{
	var manifest ReleaseManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
      return nil, err
	}
	
	fields := map[string]interface{}{
 		"image_url":  manifest.ImageUrl,
        "image_tag":  manifest.Tag,
        "image_size": manifest.ImageSize,
        "image_sha":  manifest.ImageSha,
	}

	for name, value := range fields {
		if value == nil {
			return nil, fmt.Errorf("missing required field: %s", name)
		}
	}

	return &manifest, nil

}

func VerifyFile(path string, expectedSha string)(bool, error){
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return false, err
	}
	computedHash := hex.EncodeToString(hasher.Sum(nil))
	
	if expectedSha == computedHash {
		return true, nil
	}
	return false, fmt.Errorf("Hash mismatch, expected %s, got %s", expectedSha, computedSha)
}
func main() {
	blob := ["0.0.0","0.1.0","https://docker.io/something","SHA256:65465465469879798798"]
	var inventory []Size
	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}

	counts := make(map[Size]int)
	for _, size := range inventory {
		counts[size] += 1
	}

	fmt.Printf("Inventory Counts:\n* Small:        %d\n* Large:        %d\n* Unrecognized: %d\n",
		counts[Small], counts[Large], counts[Unrecognized])

}

