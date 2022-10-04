package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nodecord",
	Short: "Doğrulayıcı takip ve alarm sistemi",
	Long: `Cosmos SDK tabanlı ağlarda alarm gorevi gorur. by Nodeist
	
Sunları kontrol eder:
- Slashin periyodu
- Validatorun kaçırdıgı bloklar
- Jail olma durumu
- Tombstone olma durumu

Discord sunucusuna asagıdaki durumları rapor eder:
- Validatorun anlık durumu
- Sorun tespiti alarmı
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
