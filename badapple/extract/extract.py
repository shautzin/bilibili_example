import cv2 as cv

'''
抽取视频帧图像像素，导出成 60x45 的像素数据
'''

FRAME_FOLDER = "frames"  # 帧图像文件夹
DEPTH_45P_FOLDER = "depth45p"  # 像素文件夹
BLOCK = 24  # 24x24 = 1 个块


## 读视频
def extract():
    vc = cv.VideoCapture('D:/Workspace/Bilibili/2019-08-13-badapple/BadApple.mp4')
    c = 1

    if vc.isOpened():
        rval, frame = vc.read()

    while rval:  # 循环读取视频帧
        rval, frame = vc.read()
        cv.imwrite('%s/%d.jpg' % (FRAME_FOLDER, c), frame)
        read_pixel(c)
        c = c + 1
        cv.waitKey(1)
    vc.release()


## 读取图像文件
def read_pixel(c):
    img = cv.imread("%s/%d.jpg" % (FRAME_FOLDER, c))
    img = cv.cvtColor(img, cv.COLOR_BGR2RGB)  # 转为 RGB 顺序
    (height, width, _) = img.shape

    depth_file = open('%s/%d.txt' % (DEPTH_45P_FOLDER, c), mode='w')

    # 逐块读取图像
    for v_block in range(0, height // BLOCK):
        for h_block in range(0, width // BLOCK):
            depth_file.write("%d," % read_block(img, v_block, h_block))
        depth_file.write("\n")
    depth_file.close()


## 计算 整块 像素值平均值
def read_block(img, v_block, h_block):
    rgb_all = 0
    for h in range(v_block * BLOCK, (v_block + 1) * BLOCK):
        for w in range(h_block * BLOCK, (h_block + 1) * BLOCK):
            pixel = img[h, w]
            rgb_all += (int(pixel[0]) + int(pixel[1]) + int(pixel[2]))
    return rgb_all // (BLOCK * BLOCK * 3)


if __name__ == '__main__':
    extract()
