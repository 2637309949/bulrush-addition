/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush Limit plugin]
 */

package addition

import "github.com/2637309949/bulrush-addition/logger"

// RushLogger just for console log
var RushLogger = logger.CreateLogger(logger.SILLYLevel, nil,
	[]*logger.Transport{
		&logger.Transport{
			Level: logger.SILLYLevel,
		},
	},
)
