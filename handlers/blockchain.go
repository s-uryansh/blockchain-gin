package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s-uryansh/blockchain-gin/configs"
	"github.com/s-uryansh/blockchain-gin/models"
)

// check if there are existing blocks in the blockchain and returh ok status
func GetBlockchain(c *gin.Context) {
	var blockchain []models.Block
	configs.DB.Find(&blockchain)
	c.JSON(http.StatusOK, blockchain)
}

func AddBlock(c *gin.Context) {
	var blockchain []models.Block
	configs.DB.Order("id desc").Limit(1).Find(&blockchain)

	if len(blockchain) == 0 {
		var newData struct {
			Data string `json:data`
		}
		if c.BindJSON(&newData) == nil {
			newBlock := models.Block{
				Index:        0,
				Timestamp:    time.Now().Format("01/02 03:04:05PM '06 MST"),
				PreviousHash: "0",
				Data:         newData.Data,
			}
			// newBlock.Hash = models.CalculateHash(newBlock)
			models.ProofOfWork(&newBlock)

			configs.DB.Create(&newBlock)
			c.JSON(http.StatusOK, gin.H{
				"Server": "First block added",
				"Block":  newBlock,
			})
		}
		return
	}

	prevBlock := blockchain[0] //saving prev block
	// fmt.Println(prevBlock.Index)

	var newData struct {
		Data string `json:data`
	}

	if c.BindJSON(&newData) == nil {
		newBlock := models.Block{
			Index:        prevBlock.Index + 1,
			Timestamp:    time.Now().String(),
			PreviousHash: prevBlock.Hash,
			Data:         newData.Data,
		}
		// newBlock.Hash = models.CalculateHash(newBlock)
		models.ProofOfWork(&newBlock)

		configs.DB.Create(&newBlock)
		c.JSON(http.StatusOK, gin.H{
			"Server": "Block added",
			"Block":  newBlock,
		})
	}
}
