package pkg

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"
	"wucms-gva/server/global"
	"wucms-gva/server/model/common/request"
	"wucms-gva/server/model/common/response"
	"wucms-gva/server/model/pkg"
	"wucms-gva/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type Reg struct{}

func (r *Reg) CreateRedis(c *gin.Context) {
	var ctx = context.Background()
	// 连接Redis数据库
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 测试连接Redis
	ping, _ := client.Ping(ctx).Result()
	fmt.Printf("测试连接Redis：%v\n", ping)
	defer client.Close()
	// 设置字符串类型的数据
	// 参数ctx是内置包context创建的上下文对象
	// 参数key和value是键值对，数据类型为字符串
	// 参数expiration是有效期，数据类型为time.Duration
	strSet, _ := client.Set(ctx, "name", "Tom", time.Hour).Result()
	fmt.Printf("设置字符串类型的数据：%v\n", strSet)
	// 获取字符串类型的数据
	strGet, _ := client.Get(ctx, "name").Result()
	fmt.Printf("获取字符串类型的数据：%v\n", strGet)
	// 删除字符串类型的数据
	// 参数keys是不固定参数，参数类型为字符串，代表字符串值
	strDel, _ := client.Del(ctx, "name").Result()
	fmt.Printf("删除字符串类型的数据：%v\n", strDel)
	// 设置哈希类型的数据
	// 参数ctx是内置包context创建的上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数values是不固定参数，参数类型为空接口，代表哈希数值
	hashHset, _ := client.HSet(ctx, "Tom", "age", 10).Result()
	fmt.Printf("设置哈希类型的数据：%v\n", hashHset)
	// 获取哈希类型的数据
	// 参数field是值，数据类型为字符串类型
	hashHGet, _ := client.HGet(ctx, "Tom", "age").Result()
	fmt.Printf("获取哈希类型的数据：%v\n", hashHGet)
	// 删除哈希类型的数据
	// 参数fields是不固定参数，数据类型为字符串类型，代表哈希数值
	hashHDel, _ := client.HDel(ctx, "Tom", "age").Result()
	fmt.Printf("删除哈希类型的数据：%v\n", hashHDel)
	// 在列表中添加一个或多个值
	// 参数ctx是内置包context创建的上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数values是不固定参数，参数类型为空接口，代表列表元素
	litRPush, _ := client.RPush(ctx, "Tom", "English", "Chinese").Result()
	fmt.Printf("在列表中添加一个或多个值：%v\n", litRPush)
	// 获取列表指定范围内的元素
	// 参数start和stop是列表索引，数据类型为整型
	litLRange, _ := client.LRange(ctx, "Tom", 0, 2).Result()
	fmt.Printf("获取列表指定范围内的元素：%v\n", litLRange)
	// 移出并获取列表的第一个元素
	// 参数timeout设置超时，数据类型为time.Duration
	// 参数keys是不固定参数，参数类型为字符串，代表列表元素
	litBLPop, _ := client.BLPop(ctx, time.Second, "Tom").Result()
	fmt.Printf("移出并获取列表的第一个元素：%v\n", litBLPop)
	// 向集合添加一个或多个成员
	// 参数ctx是内置包context创建的上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数members是不固定参数，参数类型为空接口，代表集合成员值
	SetSadd, _ := client.SAdd(ctx, "Tim", 20, "170CM").Result()
	fmt.Printf("向集合添加一个或多个成员：%v\n", SetSadd)
	// 获取集合中的所有成员
	SetSMembers, _ := client.SMembers(ctx, "Tim").Result()
	fmt.Printf("向集合添加一个或多个成员：%v\n", SetSMembers)
	// 移除并返回集合中的一个随机元素
	SetSPop, _ := client.SPop(ctx, "Tim").Result()
	fmt.Printf("移除并返回集合中的一个随机元素：%v\n", SetSPop)
	// 有序集合添加或更新一个或多个成员和成员的分数
	// 参数ctx是内置包context创建的上下文对象
	// 参数key是键，数据类型为字符串类型
	// 参数members是不固定参数，数据类型是结构体Z的实例化对象，包含集合成员和分数
	z1 := redis.Z{Member: "170CM", Score: 5}
	z2 := redis.Z{Member: 10, Score: 10}
	ZsetZAdd, _ := client.ZAdd(ctx, "Tim", &z1, &z2).Result()
	fmt.Printf("移除并返回集合中的一个随机元素：%v\n", ZsetZAdd)
	// 通过索引区间返回有序集合指定区间内的成员
	// 参数start和stop是有序集合的索引区间，数据类型为整型
	ZsetZRange, _ := client.ZRange(ctx, "Tim", 0, 2).Result()
	fmt.Printf("通过索引区间返回有序集合指定区间内的成员：%v\n", ZsetZRange)
	// 移除有序集合中的一个或多个成员
	ZsetZRem, _ := client.ZRem(ctx, "Tim", z1).Result()
	fmt.Printf("移除有序集合中的一个或多个成员：%v\n", ZsetZRem)
	// 新增流类型数据
	// 参数ctx是内置包context创建的上下文对象
	// 参数XAddArgs是指针类型的结构体XAddArgs的实例化对象，代表流类型的数据结构
	// 实例化结构体XAddArgs只需设置成员Stream和Values
	x1 := redis.XAddArgs{
		Stream: "Lily",
		Values: map[string]interface{}{"age": 10, "height": "160CM"},
	}
	// 结构体XAddArgs实例化对象以指针方式作为参数
	streXAdd, _ := client.XAdd(ctx, &x1).Result()
	fmt.Printf("新增流类型数据：%v\n", streXAdd)
	// 获取流类型所有数据
	// 参数stream代表流数据名称，即结构体XAddArgs的成员Stream
	// 参数start和stop是最小值和最大值，以“-”和“+”表示
	streXRange, _ := client.XRange(ctx, "Lily", "-", "+").Result()
	fmt.Printf("获取流类型所有数据：%v\n", streXRange)
	// 遍历变量streXRange，遍历对象为结构体XMessage，结构体成员ID是流数据ID
	for _, v := range streXRange {
		fmt.Printf("获取流类型所有数据的ID：%v\n", v.ID)
		// 通过流数据ID删除数据
		streXDel, _ := client.XDel(ctx, "Lily", v.ID).Result()
		fmt.Printf("ID：%v已删除，数据量：%v\n", v.ID, streXDel)
	}
	// 新增字符串类型数据
	client.Set(ctx, "Tim", "ABCDEFGHIJKLMN", 0)
	// 将字符串类型数据转为二进制数据，然后修改二级制的位数
	// 参数key代表redis的键
	// 参数offset是二级制的位数偏移量，0代表从左边第一位算起
	// 参数value只有0和1，因为二级制只有0和1
	bitSetBit, _ := client.SetBit(ctx, "Tim", 0, 1).Result()
	fmt.Printf("位图类型数据：%v\n", bitSetBit)
	// 获取位图类型数据某个偏移量的值
	bitGetBit, _ := client.GetBit(ctx, "Tim", 0).Result()
	fmt.Printf("获取位图类型数据某个偏移量的值：%v\n", bitGetBit)
	// 删除位图数据，即删除字符串数据
	bitDel, _ := client.Del(ctx, "Tim").Result()
	fmt.Printf("删除位图数据：%v\n", bitDel)
}

func (r *Reg) CreateEx(c *gin.Context) {
	fmt.Println("9999999999999999")
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 10000)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	// f, err := excelize.OpenFile("Book1.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer func() {
	// 	// Close the spreadsheet.
	// 	if err := f.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// // Get value from cell by given worksheet name and cell reference.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	// // Get all the rows in the Sheet1.
	// rows, err := f.GetRows("Sheet1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

}

func (r *Reg) CreateReg(c *gin.Context) {
	var Reg pkg.Reg
	err := c.ShouldBindJSON(&Reg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	matched, _ := regexp.MatchString("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", strconv.Itoa(Reg.Phone))
	if !matched {
		response.FailWithMessage("手机号码有误", c)
		return
	}
	Reg.RegTime, err = utils.ParseAndFormatTime(Reg.Time)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Create(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{}, "创建成功", c)
}

func (r *Reg) DeleteReg(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBind(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var Reg pkg.Reg
	err = global.GVA_DB.Where("id = ?", request.ID).Delete(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (r *Reg) DeleteRegByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Delete(&pkg.Reg{}, "id in ?", IDS.Ids).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (r *Reg) UpdateReg(c *gin.Context) {
	var Reg pkg.Reg
	err := c.ShouldBindJSON(&Reg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Save(&Reg).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (r *Reg) FindReg(c *gin.Context) {
	var request request.GetById
	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("pppppppppppppppppppppppp")
	fmt.Println(request.ID)
	var reg pkg.Reg
	err = global.GVA_DB.Where("id = ?", request.ID).Preload("Doctor").First(&reg).Error
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"Reg": reg}, c)
	}
}

func (r *Reg) GetRegList(c *gin.Context) {
	var pageInfo pkg.RegSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&pkg.Reg{})
	var Regs []pkg.Reg
	var total int64

	if len(pageInfo.Keyword) > 0 {
		db = db.Where("name like ?", "%"+pageInfo.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if pageInfo.StartCreatedAt != nil && pageInfo.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", pageInfo.StartCreatedAt, pageInfo.EndCreatedAt)
	}
	err = db.Limit(limit).Offset(offset).Preload("Doctor").Preload("User").Order("updated_at desc").Find(&Regs).Error

	// err = global.GVA_DB.Preload("User").Preload("TermTaxonomy.Term").Find(&Regs).Error

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     Regs,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
	// response.OkWithDetailed(gin.H{"list": Regs}, "查询成功", c)
}
