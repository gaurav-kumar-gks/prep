from re import A
import numpy as np

a = np.arange(5)
a1 = np.array([1, 2, 3])
a2 = np.array([[1, 2], [3, 4]])
a3 = np.zeros((3, 4), dtype=np.int16)
a4 = np.ones(1, dtype=np.int16)
a5 = np.arange(10, 30, 5)
a6 = np.linspace(0, 1, 10)
a7 = np.linspace(0, 1, 5)
# print(a)
# print(type(a))
# print(a.dtype)
# print(a.shape)
# print(a.itemsize)
# print(a.ndim)
# print(a.size)
# print(a * 2)
# print(a ** 2)
# print(a + 1)
# print(a - 1)
# print(a.sum())
# print(a.min())
# print(a.max())
# print(a.prod())
# print(a.mean())
# print(a5.reshape(2,2))
# print(a6.reshape(2,5))

# a.ravel()
# a.T
aa = np.array([[1, 2], [3, 4]])
ab = np.array([[-1, -2], [-3, -4]])
# print("aa: ", aa)
# print("ab: ", ab)
# print("vst: ", np.vstack((aa, ab)))
# print("hst: ", np.hstack((aa, ab)))
# print("cst: ", np.column_stack((aa, ab)))
# print("hsp: ", np.hsplit(aa, 2))
# print("hsp: ", np.vsplit(aa, 2))


# np.sort()
# np.argsort()
# np.partition()
# np.concatenate((a, b))
# np.unique(a)
# np.flip(a)

# a[np.newaxis, :] # row vector
# a[:, np.newaxis] # col vector
# np.expand_dims(a, axis=1)
# a[(a > 5) | (a < 2)]

# a[np.nonzero(a < 5)]
